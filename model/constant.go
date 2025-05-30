package model

import "time"

var (
	LogLevel        = "log-level"
	LogLevelInfo    = "info"
	LogLevelError   = "error"
	LogLevelDebug   = "debug"
	LogLevelWarning = "warn"
)

var (
	ApiPackage        = "api"
	StorePackage      = "store"
	ControllerPackage = "controller"
	ModelPackage      = "model"
	UtilPackage       = "util"
	MainPackage       = "main"
)
var (
	Controller = "controller"
	Store      = "store"
	Api        = "api"
	Main       = "main"
)
var (
	NewServer       = "newServer"
	NewStore        = "newStore"
	CreateUser      = "create-User"
	GetUsers        = "get-Users"
	GetUser         = "get-User"
	SignUP          = "sign-up"
	SignIn          = "sign-in"
	GetUserByFilter = "get-user-by-filter"
	UpdateUser      = "update-user"
	DeleteUser      = "delete-user"

	CreateSchool      = "create-schools"
	GetSchools        = "get-schools"
	GetSchool         = "get-school"
	GetSchoolByFilter = "get-school-by-filter"
	UpdateSchool      = "update-school"
	DeleteSchool      = "delete-school"

	CreateClass      = "create-classs"
	GetClasss        = "get-classs"
	GetClass         = "get-class"
	GetClassByFilter = "get-class-by-filter"
	UpdateClass      = "update-class"
	DeleteClass      = "delete-class"

	CreateRoom      = "create-room"
	GetRoom         = "get-room"
	GetRooms        = "get-rooms"
	GetRoomByFilter = "get-room-by-filter"
	UpdateRoom      = "update-room"
	DeleteRoom      = "delete-room"

	CreateTeacher      = "create-teacher"
	GetTeacher         = "get-teacher"
	GetTeachers        = "get-teachers"
	GetTeacherByFilter = "get-teacher-by-filter"
	UpdateTeacher      = "update-teacher"
	DeleteTeacher      = "delete-teacher"

	CreatePublisher      = "create-publisher"
	GetPublisher         = "get-publisher"
	GetPublishers        = "get-publishers"
	GetPublisherByFilter = "get-publisher-by-filter"
	UpdatePublisher      = "update-publisher"
	DeletePublisher      = "delete-publisher"

	CreateAuthor      = "create-author"
	GetAuthor         = "get-author"
	GetAuthors        = "get-authors"
	GetAuthorByFilter = "get-author-by-filter"
	UpdateAuthor      = "update-author"
	DeleteAuthor      = "delete-author"

	CreateBook      = "create-book"
	GetBook         = "get-book"
	GetBooks        = "get-books"
	GetBookByFilter = "get-book-by-filter"
	UpdateBook      = "update-book"
	DeleteBook      = "delete-book"

	CreateLab      = "create-lab"
	GetLab         = "get-lab"
	GetLabs        = "get-labs"
	GetLabByFilter = "get-lab-by-filter"
	UpdateLab      = "update-lab"
	DeleteLab      = "delete-lab"

	AuthMiddlewareComplete = "AuthMiddlewareComplete"
	SetLimitAndPage        = "setLimitAndPag e"
	SetDateRangeFilter     = "setDateRangeFilter"
)

// General
var (
	Value    = "value"
	Email    = "email"
	Password = "password"
	UserID   = "UserID"
	Expire   = "exp"

	Authorization = "X-Token"

	DNS = "host=localhost user=aman password=aman123 dbname=management port=5432 sslmode=disable"

	DataPerPage = "limit"
	PageNumber  = "page"
	StartDate   = "start_date"
	EndDate     = "end_date"
	TimeLayout  = "2006-01-02 15:04:05.000 -0700"
)

var (
	TokenExpiration = time.Hour * 24
)

var SecretKey = []byte("managment-secreat-key")

// schools type
var (
	SuperAdminUser = "superAdmin"
	AdminUser      = "Admin"
	NormalUser     = "user"
	UserTypes      = []string{"superAdmin", "Admin", "user"}
)
