package main

import (
	"echo/api/routers"
	middleware "echo/middleware"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	_ "echo/cmd/api/docs"

	logger "github.com/labstack/gommon/log"

	// migrate "echo/database/connection"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	error_ := godotenv.Load()
	if error_ != nil {
		log.Fatal("Error", error_)
	}
	// migrate.HandleMigration()
	// migrate.Generate()

}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	router := echo.New()
	router.Debug = true
	if l, ok := router.Logger.(*logger.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
	}

	// DocsGenerate()
	router.GET("/swagger/*", echoSwagger.WrapHandler)
	middleware.MiddlewareHandler(router)
	routers.Routers(router)

	severPort := os.Getenv("PORT")
	router.Logger.Fatal(router.Start(":" + severPort))
}
