package main

import (
	"fmt"

	"github.com/Aman17101/SchoolMangement/api"
	"github.com/Aman17101/SchoolMangement/controller"
)

func main() {

	api := api.ApiRouts{}
	api.StartApp(controller.Server{})

	fmt.Printf(" main sever =%v\n", api)
}
