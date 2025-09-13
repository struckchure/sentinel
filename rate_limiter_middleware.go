package sentinel

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func RateLimiterMiddleware(backend Backend, config map[string]any) echo.MiddlewareFunc {
	_limit, _ := config["limit"].(int)
	_burst, _ := config["burst"].(int)
	_expires, _ := config["expires"].(string)

	expiresIn, err := time.ParseDuration(_expires)
	if err != nil {
		logger.Warn(err.Error())
		return nil
	}

	rateLimiterConfig := middleware.RateLimiterConfig{
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{
				Rate:      rate.Limit(_limit),
				Burst:     _burst,
				ExpiresIn: expiresIn,
			},
		),
	}

	return middleware.RateLimiterWithConfig(rateLimiterConfig)
}
