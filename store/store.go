package store

import (
	"fmt"

	"github.com/Aman17101/SchoolMangement/model"
	"github.com/Aman17101/SchoolMangement/util"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgress struct {
	DB *gorm.DB
}

func (store *Postgress) NewStore() error {
	
	util.Log(model.LogLevelInfo, model.StorePackage, model.NewStore, "creating new store", nil)
	db, err := gorm.Open(postgres.Open(model.DNS), &gorm.Config{})
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, " err while creating new store", err)
		return err
	} else {
		store.DB = db
	}
	err = db.AutoMigrate(
		model.User{},
		model.School{},
	)

	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "err while running automigration", err)
		return err
	}

	fmt.Printf("db =%v\n", db)
	return nil

}

type StoreOperation interface {
	NewStore() error
	CreateUser(user *model.User) error
	GetUsers() ([]model.User, error)
	GetUser(uuid.UUID) (model.User, error)
	SingIn(userSignIn model.UserSignIn) (*model.User, error)
	 SignUp(user *model.User) error

	 CreateSchool(school *model.School) error
	GetSchools() ([]model.School, error)
	GetSchool(uuid.UUID) (model.School, error)
	
	
}