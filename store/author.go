package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/Aman17101/SchoolMangement/model"
	"github.com/Aman17101/SchoolMangement/util"
	"gorm.io/gorm"
)

func (store Postgress) CreateAuthor(author *model.Author) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateAuthor, "creating new author", nil)
	response := store.DB.Create(author)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "error while creating new author", response.Error)
		return response.Error
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateAuthor, "Created new author", author)
	return nil
}

func (store Postgress) GetAuthors() ([]model.Author, error) {

	authors := []model.Author{}
	util.Log(model.LogLevelInfo, model.StorePackage, model.GetAuthors, "fetching records of author from db", nil)
	if err := store.DB.Find(&authors).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetAuthors, "error while fetching authors from db", err)
		return authors, err
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetAuthors, "records of author from db", authors)
	return authors, nil
}

func (store Postgress) GetAuthor(authorID uuid.UUID) (model.Author, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetAuthor, "fetching records of author from db", nil)
	var author model.Author
	if err := store.DB.First(&author, "id = ?", authorID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			util.Log(model.LogLevelError, model.StorePackage, model.GetAuthor, "author record not found", err)
		} else {
			util.Log(model.LogLevelError, model.StorePackage, model.GetAuthor, "error while fetching author from db", err)
		}
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetAuthor, "records of author from db", author)
	return author, nil
}

func (store Postgress) GetAuthorByFilter(filter map[string]interface{}) ([]model.Author, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetAuthorByFilter, "fetching records of author from db", nil)
	var authors []model.Author
	query := store.DB

	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == model.StartDate || key == model.EndDate {
			continue
		}
		util.Log(model.LogLevelInfo, model.StorePackage, model.GetAuthorByFilter,
			"filters key", key+" value = "+fmt.Sprintf("%v", value))
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	setLimitAndPage(filter, query)
	setDateRangeFilter(filter, query)

	err := query.Find(&authors).Error
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetAuthorByFilter,
			"error while reading author data", err)
		return nil, fmt.Errorf("error while fetching author records from DB, err = %v", err)
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetAuthorByFilter, "records of authors from db", authors)
	return authors, nil
}

func (store Postgress) UpdateAuthor(author *model.Author) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateAuthor, "updating author data", *author)
	resp := store.DB.Save(author)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.UpdateAuthor,
			"error while updating author data", resp.Error)
		return fmt.Errorf("error while updating author record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateAuthor,
		"successfully updated author", nil)
	return nil
}

// DeleteAuthor is used to delete record by given authorID
func (store Postgress) DeleteAuthor(authorID string) error {

	var author model.Author
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteAuthor, "deleting author data", authorID)
	if err := store.DB.First(&author, `"id" = '`+authorID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.DeleteAuthor,
			"error while deleting author data", err)
		return fmt.Errorf("author not found for given id, ID = %v", authorID)
	}
	resp := store.DB.Delete(author)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting author record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteAuthor,
		"successfully deleted author", nil)
	return nil
}

