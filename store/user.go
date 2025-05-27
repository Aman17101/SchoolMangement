package store

import (
	"github.com/Aman17101/SchoolMangement/model"
	"github.com/Aman17101/SchoolMangement/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (store Postgress) CreateUser(user *model.User) error {
	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateUser, " creating new user", nil)
	response := store.DB.Create(user)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.CreateUser, "err while creating new user", response.Error)
		return response.Error
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateUser, "successfully created user ", user)
	return nil
}

//GET request for all the data

func (store Postgress) GetUsers() ([]model.User, error) {
	users := []model.User{}
	util.Log(model.LogLevelInfo, model.StorePackage, model.GetUsers, " fetching record for user from db ", nil)

	if err := store.DB.Find(&users).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetUsers, "err while fetching users from db", err)
		return users, err
	}
	util.Log(model.LogLevelError, model.StorePackage, model.GetUsers, "records of user from db ", users)
	return users, nil
}

// GET request on the basis of particular id

func (store Postgress) GetUser(userID uuid.UUID) (model.User, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetUser, " fetching record for user from db ", nil)
	var user model.User
	if err := store.DB.First(&user, "id =?", userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			util.Log(model.LogLevelError, model.StorePackage, model.GetUser, " user record not found ", err)
		} else {
			util.Log(model.LogLevelError, model.StorePackage, model.GetUser, " error while fetching user from db ", err)
		}
	}

	util.Log(model.LogLevelError, model.StorePackage, model.GetUser, "records of user from db ", user)
	return user, nil
}
