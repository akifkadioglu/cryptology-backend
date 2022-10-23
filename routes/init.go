package routes

import (
	"github.com/labstack/echo/v4"
	"setup/adapter"
	"setup/env"
)

var E = echo.New()
var Network = E.Group("", adapter.ConsoleAdapter)

func Init() {
	Api()
	Web()
	addr := env.GoDotEnvVariable("APP_HOST")
	E.Logger.Fatal(E.Start(addr))
}
