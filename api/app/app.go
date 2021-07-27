package app

import (
	"golang-testing/docs"
	"golang-testing/internal/pkg/config"
	"os"

	"golang-testing/api/controllers"
	"golang-testing/internal/pkg/httpserver"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
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
	confSwagger()
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

	//Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

func confSwagger() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "APP usin GetCountry from an external API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3001"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
