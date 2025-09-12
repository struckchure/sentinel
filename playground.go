package sentinel

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samber/lo"
)

func Play() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	url1 := lo.Must(url.Parse("http://localhost:8010"))
	url2 := lo.Must(url.Parse("http://localhost:8020"))
	url3 := lo.Must(url.Parse("http://localhost:8030"))

	lb := middleware.ProxyWithConfig(middleware.ProxyConfig{
		Rewrite: map[string]string{
			"^/todos/":  "/todos/",
			"^/todos/*": "/todos/$1",
		},
		Balancer: middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
			{URL: url1},
			{URL: url2},
			{URL: url3},
		}),
		ModifyResponse: func(r *http.Response) error {
			r.Header.Set("Server", "Sentinel")
			return nil
		},
	})

	g1 := e.Group("/todos")
	g1.Use(lb)
	g2 := e.Group("/todos/*")
	g2.Use(lb)

	e.Logger.Fatal(e.Start(":8000"))
}
