package sentinel

import (
	"context"
	"fmt"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type AuthNJwtMiddlewareConfig struct {
	Alg             *string `mapstructure:"alg"`
	JwksUrl         *string `mapstructure:"jwks_url"`
	JwtSecret       *string `mapstructure:"jwt_secret"`
	PropagateClaims []struct {
		From string `mapstructure:"from"`
		To   string `mapstructure:"to"`
	} `mapstructure:"propagate_claims"`
}

func AuthNJwtMiddleware(backend Backend, _config map[string]any) echo.MiddlewareFunc {
	var config AuthNJwtMiddlewareConfig

	err := mapstructure.Decode(_config, &config)
	if err != nil {
		logger.Warn(err.Error())
		return nil
	}

	jwtMiddlewareConfig := echojwt.Config{}

	if config.JwksUrl != nil {
		k, err := keyfunc.NewDefaultCtx(context.Background(), []string{*config.JwksUrl})
		if err != nil {
			err = fmt.Errorf("failed to create a keyfunc.Keyfunc from the server's URL.\nError: %s", err)
			logger.Warn(err.Error())
			return nil
		}

		jwtMiddlewareConfig.KeyFunc = k.Keyfunc
	}

	if config.Alg != nil {
		jwtMiddlewareConfig.SigningMethod = *config.Alg
	}

	if config.JwtSecret != nil && config.JwksUrl == nil {
		fmt.Println("setting sign key")
		jwtMiddlewareConfig.SigningKey = []byte(*config.JwtSecret)
	}

	jwtMiddlewareConfig.SuccessHandler = func(c echo.Context) {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		for _, propagateClaim := range config.PropagateClaims {
			c.Response().Header().Set(propagateClaim.To, fmt.Sprintf("%s", claims[propagateClaim.From]))
		}
	}

	return echojwt.WithConfig(jwtMiddlewareConfig)
}
