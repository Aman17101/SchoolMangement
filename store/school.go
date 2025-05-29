package store

import (
	

	"github.com/Aman17101/SchoolMangement/model"
	"github.com/Aman17101/SchoolMangement/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)
//POST request to create data
func (store Postgress) CreateSchool(school *model.School) error {
	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateSchool, " creating new school", nil)
	response := store.DB.Create(school)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.CreateSchool, "err while creating new school", response.Error)
		return response.Error
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateSchool, "successfully created school ", school)
	return nil
}

//GET request for all the data

func (store Postgress) GetSchools() ([]model.School, error) {
	schools := []model.School{}
	util.Log(model.LogLevelInfo, model.StorePackage, model.GetSchools, " fetching record for school from db ", nil)

	if err := store.DB.Find(&schools).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetSchools, "err while fetching schools from db", err)
		return schools, err
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.GetSchools, "records of school from db ", schools)
	return schools, nil
}

// GET request on the basis of particular id

func (store Postgress) GetSchool(schoolID uuid.UUID) (model.School, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetSchool, " fetching record for school from db ", nil)
	var school model.School
	if err := store.DB.First(&school, "id =?", schoolID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			util.Log(model.LogLevelError, model.StorePackage, model.GetSchool, " school record not found ", err)
		} else {
			util.Log(model.LogLevelError, model.StorePackage, model.GetSchool, " error while fetching school from db ", err)
		}
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetSchool, "records of school from db ", school)
	return school, nil
}

