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
		model.Class{},
		model.Room{},
		model.Book{},
		model.Teacher{},
		model.Author{},
		model.Publisher{},
		model.Lab{},
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
	GetUserByFilter(filter map[string]interface{}) ([]model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(userID string) error

	CreateSchool(school *model.School) error
	GetSchools() ([]model.School, error)
	GetSchool(uuid.UUID) (model.School, error)
	GetSchoolByFilter(filter map[string]interface{}) ([]model.School, error)
	UpdateSchool(school *model.School) error
	DeleteSchool(schoolID string) error

	CreateClass(class *model.Class) error
	GetClasss() ([]model.Class, error)
	GetClass(uuid.UUID) (model.Class, error)
	GetClassByFilter(filter map[string]interface{}) ([]model.Class, error)
	UpdateClass(class *model.Class) error
	DeleteClass(classID string) error

	CreateTeacher(teacher *model.Teacher) error
	GetTeachers() ([]model.Teacher, error)
	GetTeacher(uuid.UUID) (model.Teacher, error)
	GetTeacherByFilter(filter map[string]interface{}) ([]model.Teacher, error)
	UpdateTeacher(teacher *model.Teacher) error
	DeleteTeacher(teacherID string) error

	CreateRoom(room *model.Room) error
	GetRooms() ([]model.Room, error)
	GetRoom(uuid.UUID) (model.Room, error)
	GetRoomByFilter(filter map[string]interface{}) ([]model.Room, error)
	UpdateRoom(room *model.Room) error
	DeleteRoom(roomID string) error
	

	CreateAuthor(author *model.Author) error
	GetAuthors() ([]model.Author, error)
	GetAuthor(uuid.UUID) (model.Author, error)
	GetAuthorByFilter(filter map[string]interface{}) ([]model.Author, error)
	UpdateAuthor(author *model.Author) error
	DeleteAuthor(authorID string) error

	CreatePublisher(publisher *model.Publisher) error
	GetPublishers() ([]model.Publisher, error)
	GetPublisher(uuid.UUID) (model.Publisher, error)
	GetPublisherByFilter(filter map[string]interface{}) ([]model.Publisher, error)
	UpdatePublisher(publisher *model.Publisher) error
	DeletePublisher(publisherID string) error

	CreateBook(book *model.Book) error
	GetBooks() ([]model.Book, error)
	GetBook(uuid.UUID) (model.Book, error)
	GetBookByFilter(filter map[string]interface{}) ([]model.Book, error)
	UpdateBook(book *model.Book) error
	DeleteBook(bookID string) error

	CreateLab(lab *model.Lab) error
	GetLabs() ([]model.Lab, error)
	GetLab(uuid.UUID) (model.Lab, error)
	GetLabByFilter(filter map[string]interface{}) ([]model.Lab, error)
	UpdateLab(lab *model.Lab) error
	DeleteLab(labID string) error
}
