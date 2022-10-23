package routes

import (
	"setup/adapter"
	"setup/controllers"
	authcontroller "setup/controllers/AuthController"
	recordscontroller "setup/controllers/RecordsController"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Web() {

	Network.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	//Authenticate
	web := Network.Group("")

	web.POST("/register", authcontroller.Register)
	web.POST("/login", authcontroller.Login)

	//For Auth Users

	auth := web.Group("")
	auth.Use(adapter.AuthAdapter())
	auth.GET("/getRecord", recordscontroller.GetRecord)
	auth.POST("/createRecord", recordscontroller.CreateRecord)

	//index
	Network.GET("/index", controllers.Index)

}
