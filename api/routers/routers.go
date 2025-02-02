// This package is responsible for route the request to the appropriate request
// handler functions.
package routers

import (
	"github.com/labstack/echo/v4"
)

func Routers(r *echo.Echo) {

	// `v1` is a variable that is used to group routes under the version "v1". It
	// helps in organizing and grouping related routes together for better
	// maintainability and readability of the code.
	version_v1 := r.Group("/v1")

	v1_authRouteGroup := version_v1.Group("/auth")
	AuthRoutes(v1_authRouteGroup)

	v1_userRouteGroup := version_v1.Group("/users")
	UsersRoutes(v1_userRouteGroup)

}
