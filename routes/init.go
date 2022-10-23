package routes

import (
	"github.com/labstack/echo/v4"
	"setup/adapter"
)

var E = echo.New()
var Network = E.Group("", adapter.ConsoleAdapter)

func Init() {
	Api()
	Web()
	addr := "https://cryptology-homework.herokuapp.com"
	E.Logger.Fatal(E.Start(addr))
}
