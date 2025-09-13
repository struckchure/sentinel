package sentinel

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func AllowedMethodMiddleware(backend Backend, config any) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !lo.Contains(backend.Methods, Method(c.Request().Method)) {
				return c.JSON(http.StatusMethodNotAllowed, echo.Map{"message": "method not allowed"})
			}

			if err := next(c); err != nil {
				c.Error(err)
			}

			return nil
		}
	}
}
