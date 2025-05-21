package controller

import (
	"fmt"

	"github.com/Aman17101/SchoolMangement/store"
)

type Server struct {
	PostgressDb store.StoreOperation
}

func (s *Server) NewServer(pgstore store.Postgress) {
	s.PostgressDb = &pgstore

	s.PostgressDb.NewStore()
	fmt.Printf("server =%v\n", s)
}
type ServerOperations interface{
	NewServer(pgstore store.Postgress) 
}