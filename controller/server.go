package controller

import (
	"log"

	"github.com/Aman17101/SchoolMangement/model"
	"github.com/Aman17101/SchoolMangement/store"
	"github.com/Aman17101/SchoolMangement/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	PostgressDb store.StoreOperation
}

func (s *Server) NewServer(pgstore store.Postgress) {
	util.SetLogger()
	util.Logger.Infof("Logger setup done ...\n")

	s.PostgressDb = &pgstore

	err := s.PostgressDb.NewStore()
	if err != nil {
		util.Logger.Errorf("error while creating store connection ,err=%v", err)
		util.Log(model.LogLevelError, model.Controller, "NewServer", "Error while creating store connection", err)

	} else {
		util.Logger.Infof("store connection created successfully\n")
		util.Log(model.LogLevelInfo, model.Controller, model.NewServer, "Connected with store", nil)

	}

	log.Printf("server =%v\n", s)
}

type ServerOperations interface {
	NewServer(pgstore store.Postgress)
	CreateUser(ctx *gin.Context)
}
