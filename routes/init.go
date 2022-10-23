package routes

import (
	"os"
	"setup/adapter"

	"github.com/labstack/echo/v4"
)

var E = echo.New()
var Network = E.Group("", adapter.ConsoleAdapter)

func Init() {
	Api()
	Web()
	port := "0.0.0.0:" + os.Getenv("PORT")
	if port == "" {
		port = "0.0.0.0:9000" // Default port if not specified
	}
	E.Logger.Fatal(E.Start(port))
}
