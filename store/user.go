package store

import (
	"github.com/Aman17101/SchoolMangement/model"
	"github.com/Aman17101/SchoolMangement/util"
)

func (store Postgress) CreateUser(user *model.User) error {
	util.Log(model.LogLevelInfo, model.StorePackage, model.NewStore, " creating new store", nil)
	response := store.DB.Create(user)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "err while creating new store", response.Error)
		return response.Error
	}
	util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "err while creating new store", user)
	return nil
}
