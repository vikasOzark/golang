package routers

import (
	users "echo/internal/handlers"

	"github.com/labstack/echo/v4"
)

func UsersRoutes(authRouteGroup *echo.Group) {
	authRouteGroup.GET("", users.Users)
}
