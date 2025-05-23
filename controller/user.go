package controller

import (
	"net/http"
	"time"

	"github.com/Aman17101/SchoolMangement/model"
	"github.com/Aman17101/SchoolMangement/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server Server) CreateUser(ctx *gin.Context) {

	util.Log(model.LogLevelError, model.ControllerPackage, model.NewStore, "creating new store", nil)
	user := model.User{}
	err := ctx.Bind(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.NewStore, "error while json binding ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	user.ID = uuid.New()
	user.CreatedAt = time.Now()

	err = server.PostgressDb.CreateUser(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.NewStore, "error while json binding ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}
	ctx.JSON(http.StatusCreated, user)

}
