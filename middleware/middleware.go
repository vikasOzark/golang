package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var logger = "method:${method} uri:${uri} status:${status}\n"

func MiddlewareHandler(router *echo.Echo) {
	router.Use(middleware.CORS())
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: logger,
	}))

}
