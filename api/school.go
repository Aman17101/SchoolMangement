package api

import (
	"github.com/Aman17101/SchoolMangement/model"
	"github.com/Aman17101/SchoolMangement/util"
	"github.com/gin-gonic/gin"
)

func (api ApiRouts) SchoolRouts(routes *gin.Engine) {
	// Define routes
	group := routes.Group("school")
	{
		group.POST("/create", api.AuthMiddlewareComplete(), api.CreateSchool)
		group.GET("/:id", api.AuthMiddlewareComplete(), api.GetSchool)
		group.GET("/all", api.AuthMiddlewareComplete(), api.GetSchools)
	

	}

}

// Handler to create a school
// @router /school/create [post]
// @summary Create a school
// @tags schools
// @accept json
// @produce json
// @param school body model.School true "School object"
// @success 201 {object} model.School
// @Security ApiKeyAuth
func (api ApiRouts) CreateSchool(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateSchool, "creating new school", nil)
	api.Server.CreateSchool(ctx)
}

// Handler to get a school by ID
// @router /school/{id} [get]
// @summary Get a school by ID
// @tags schools
// @produce json
// @param id path string true "School ID"
// @success 200 {object} model.School
// @Security ApiKeyAuth
func (api ApiRouts) GetSchool(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetSchool, "fetching  school", nil)
	api.Server.GetSchool(ctx)
}

// Handler to get all schools
// @router /school/all [get]
// @summary Get all schools
// @tags schools
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.School
// @Security ApiKeyAuth
func (api ApiRouts) GetSchools(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetSchools, "fetching schools", nil)
	api.Server.GetSchools(ctx)
}

