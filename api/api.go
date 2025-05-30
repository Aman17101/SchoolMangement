package api

import (
	"github.com/Aman17101/SchoolMangement/controller"
	"github.com/Aman17101/SchoolMangement/store"
	"github.com/gin-gonic/gin"

	_ "github.com/Aman17101/SchoolMangement/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ApiRouts struct {
	Server controller.ServerOperations
}

func (api *ApiRouts) StartApp(routes *gin.Engine, server controller.Server) {
	api.Server = &server
	api.Server.NewServer(store.Postgress{})

	//swagger documentation
	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//user routs
	api.UserRouts(routes)


	//school routs
	api.SchoolRouts(routes)

	//class routs
	api.ClassRouts(routes)

	// teacher routs
		api.TeacherRouts(routes)

//room routs
		api.RoomRouts(routes)


		// Book routs
		api.BookRouts(routes)


		// Author routs
		api.AuthorRouts(routes)


		// publisher routs
		api.PublisherRouts(routes)


		// lab routs
		api.LabRouts(routes)


}
