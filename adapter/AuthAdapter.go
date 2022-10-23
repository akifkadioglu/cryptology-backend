package adapter

import (
	"setup/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthAdapter() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims: &models.JwtCustomClaims{},
		SigningKey: []byte("fil necati"),
	}
	return middleware.JWTWithConfig(config) 
}
