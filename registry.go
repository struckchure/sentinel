package sentinel

import "github.com/labstack/echo/v4"

type MiddlewareFunc func(Backend, map[string]any) echo.MiddlewareFunc

var MiddlewareRegistry = map[string]MiddlewareFunc{
	"rate-limiter": RateLimiterMiddleware,
	"auth-n":       AuthNJwtMiddleware,
}
