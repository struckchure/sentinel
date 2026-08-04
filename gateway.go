package sentinel

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samber/lo"
)

type IGateway interface {
	Run() error
}

type Gateway struct {
	config Config
}

func (g *Gateway) Run() error {
	e := echo.New()
	e.HideBanner = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "[${host}] ${method} ${uri} ${status}\n"}))

	rootLb := NewRootLB()
	proxyTargets := []*middleware.ProxyTarget{}

	for _, backend := range g.config.Backends {
		for _, service := range backend.Services {
			target := &middleware.ProxyTarget{URL: lo.Must(url.Parse(service.Url))}
			proxyTargets = append(proxyTargets, target)
		}

		rewrite := map[string]string{}

		for _, pattern := range backend.Patterns {
			rewrite[fmt.Sprintf("^%s", pattern.From)] = pattern.To
		}

		var lbAlgorithm middleware.ProxyBalancer

		switch backend.LoadBalancer {
		case LoadBalancerAlgorithmRoundRobin:
			lbAlgorithm = middleware.NewRoundRobinBalancer(proxyTargets)
		case LoadBalancerAlgorithmRandom:
			lbAlgorithm = middleware.NewRandomBalancer(proxyTargets)
		default:
			return errors.New("algorithm not implemented")
		}

		if backend.Hostname != nil {
			rootLb.AddLB(lbAlgorithm, []string{*backend.Hostname})
		} else {
			logger.Warn("hostname is missing or invalid")
		}

		lb := middleware.ProxyWithConfig(middleware.ProxyConfig{
			Rewrite:  rewrite,
			Balancer: middleware.ProxyBalancer(rootLb),
			ModifyResponse: func(r *http.Response) error {
				r.Header.Set("Server", "Sentinel")
				return nil
			},
		})

		middlewares := []echo.MiddlewareFunc{}
		for _, middleware := range backend.Middlewares {
			middlewareFunc, ok := MiddlewareRegistry[middleware.Name]
			if !ok {
				logger.Warn("middleware (%s) not found in registry", middleware.Name)
				continue
			}

			middlewares = append(middlewares, middlewareFunc(backend, middleware.Config))
		}

		for _, pattern := range backend.Patterns {
			group := e.Group(pattern.From)

			if len(backend.Methods) > 0 {
				group.Use(AllowedMethodMiddleware(backend, nil))
			}

			group.Use(middlewares...)
		}

		e.Use(lb)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", g.config.Host, g.config.Port)))

	return nil
}

func NewGateway(config Config) IGateway {
	return &Gateway{config: config}
}
