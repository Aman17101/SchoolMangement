package api

import (
	"github.com/Aman17101/SchoolMangement/controller"
	"github.com/Aman17101/SchoolMangement/store"
	"github.com/gin-gonic/gin"
)

type ApiRouts struct {
	Server controller.ServerOperations
}

func (api *ApiRouts) StartApp(routes *gin.Engine, server controller.Server) {
	api.Server = &server
	api.Server.NewServer(store.Postgress{})

	api.UserRouts(routes)
}
