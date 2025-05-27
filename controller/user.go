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

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreateUser, "creating new user", nil)
	user := model.User{}
	err := ctx.Bind(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateUser, "error while json binding ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	user.ID = uuid.New()
	user.CreatedAt = time.Now()

	err = server.PostgressDb.CreateUser(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateUser, "error while json binding ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, user)

}
func (server Server) GetUsers(ctx *gin.Context) {

	util.Log(model.LogLevelError, model.ControllerPackage, model.GetUsers, "fetching all user", nil)

	users, err := server.PostgressDb.GetUsers()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetUsers, "error while fetching users record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, users)

}
func (server Server) GetUser(ctx *gin.Context) {
	//id := ctx.Param("id")
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetUsers, "fetching all user ", nil)

	// Get ID from URL param
	// id := ctx.Param("id")
	// var id uint
	// if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
	// 	util.Log(model.LogLevelError, model.ControllerPackage, model.GetUsers, "invalid user ID", err)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	// 	return
	// }

	// uuid, err := uuid.Parse(id)
	// if err != nil {
	// 	util.Log(model.LogLevelError, model.ControllerPackage, model.GetUser, "invalid user ID format", err)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id format"})
	// 	return
	// }
	// Call DB layer
	user, err := server.PostgressDb.GetUsers()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetUsers, "error while fetching user", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
