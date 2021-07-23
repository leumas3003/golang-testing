package httpserver

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}

func AllowGracefulShutdown(e *echo.Echo) {

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-ch
		log.Infof("Received shutdown signal: [%v]", sig)

		log.Info("Performing graceful shutdown..")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if e != nil {
			log.Info("Shutting down HTTP server")
			if err := e.Shutdown(ctx); err != nil {
				log.Error(err)
			}
		}

		log.Info("Graceful shutdown complete.")
		os.Exit(1)
	}()
}
