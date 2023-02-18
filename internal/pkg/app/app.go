package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/o1vanova/calculator/cmd/calculator/internal/endpoint"
	"github.com/o1vanova/calculator/cmd/calculator/internal/service"
)

type App struct {
	endPoint *endpoint.EndPoint
	echoSvc  *echo.Echo
}

func New() (*App, error) {
	app := &App{}

	simpleService := service.New()
	app.endPoint = endpoint.New(simpleService)
	app.echoSvc = echo.New()

	app.echoSvc.GET("/", app.endPoint.Info)
	app.echoSvc.GET("/plus", app.endPoint.Plus)
	app.echoSvc.GET("/minus", app.endPoint.Minus)
	app.echoSvc.GET("/multify", app.endPoint.Multify)

	return app, nil
}

func (app *App) Start() error {
	err := app.echoSvc.Start(":8080")
	if err != nil {
		log.Errorf("The start of server is failed: %#v", err)
		return err
	}

	log.Info("The server is started")

	return nil
}
