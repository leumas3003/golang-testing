package app

import (
	"golang-testing/internal/pkg/config"
	"os"

	"golang-testing/api/controllers"
	"golang-testing/internal/pkg/httpserver"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Setup() *config.AppConfig {

	c, err := config.GetConfig()
	if err != nil {
		log.Errorf("Failed to open config [%v]", err)
		os.Exit(1)
	}

	/*err = db.Connect(c)
	if err != nil {
		log.Errorf("Failed to connect to postgres [%v]", err)
		os.Exit(1)
	}*/

	return c
}

func StartServer() *echo.Echo {

	c := Setup()

	e := httpserver.New()
	httpserver.AllowGracefulShutdown(e)
	setupRoutes(e)

	err := e.Start(c.HttpServer.Addr)
	if err != nil {
		//db.Cleanup()
		log.Errorf("HTTP server shutdown [%v]", err)
		os.Exit(1)
	}

	return e
}

func setupRoutes(e *echo.Echo) {
	e.GET("/locations/countries/:country_id", controllers.GetCountry)
}
