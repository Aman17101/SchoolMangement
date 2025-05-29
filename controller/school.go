package controller

import (
	"net/http"
	"time"

	"github.com/Aman17101/SchoolMangement/model"
	"github.com/Aman17101/SchoolMangement/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server Server) CreateSchool(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreateSchool, "creating new school", nil)
	school := model.School{}
	err := ctx.Bind(&school)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateSchool, "error while json binding ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	school.ID = uuid.New()
	school.CreatedAt = time.Now()

	err = server.PostgressDb.CreateSchool(&school)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateSchool, "error while json binding ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, school)

}
func (server Server) GetSchools(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetSchools, "fetching all school", nil)

	schools, err := server.PostgressDb.GetSchools()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetSchools, "error while fetching schools record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, schools)

}
func (server Server) GetSchool(ctx *gin.Context) {
	//id := ctx.Param("id")
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetSchools, "fetching all school ", nil)

	// Call DB layer
	school, err := server.PostgressDb.GetSchools()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetSchools, "error while fetching school", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
		return
	}

	ctx.JSON(http.StatusCreated, school)
}
