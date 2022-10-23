package routes

import (
	"setup/adapter"
	controller "setup/controllers"
)

func Api() {
	api := Network.Group("/api")

	
	//Auth Users
	auth := api.Group("")
	auth.Use(adapter.AuthAdapter())

	auth.GET("/index", controller.Index)
}
