package routers

import (
	auth "echo/internal/handlers"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(authRouteGroup *echo.Group) {
	login := authRouteGroup.POST("/login", auth.AuthHandler)
	login.Name = "user-login"
	authRouteGroup.POST("/registration", auth.Registration)
}
