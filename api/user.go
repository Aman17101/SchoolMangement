package api

import (
	"github.com/Aman17101/SchoolMangement/model"
	"github.com/Aman17101/SchoolMangement/util"
	"github.com/gin-gonic/gin"
)

func (api ApiRouts) UserRouts(routes *gin.Engine) {
	group := routes.Group("user")
	{
		group.POST("/create", api.CreateUser)
	}
}
func (api ApiRouts) CreateUser(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateUser, " creating new user", nil)
	api.Server.CreateUser(ctx)
}
