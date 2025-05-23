package main

import (
	"github.com/Aman17101/SchoolMangement/api"
	"github.com/Aman17101/SchoolMangement/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	api := api.ApiRouts{}
	controller := controller.Server{}
	routes := gin.Default()
	api.StartApp(routes, controller)

	routes.Run(":8080")

	//fmt.Printf(" main sever =%v\n", api)
}
