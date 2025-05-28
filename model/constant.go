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
	NewServer  = "newServer"
	NewStore   = "newStore"
	CreateUser = "create-user"
	GetUsers   = "get-users"
	GetUser    = "get-user"
	SignUP     = "sign-up"
	SignIn     = "sign-in"
	AuthMiddlewareComplete = "AuthMiddlewareComplete"
)
// General
var (
	Value    = "value"
	Email    = "email"
	Password = "password"
	UserID   = "userID"
	Expire   = "exp"

	Authorization = "X-Token"

	DNS = "host=localhost user=aman password=aman123 dbname=management port=5432 sslmode=disable"

)

var (
	TokenExpiration = time.Hour * 24
)

var SecretKey = []byte("managment-secreat-key")

// user type
var (
	
	SuperAdminUser      = "superAdmin"
	AdminUser           = "Admin"
	NormalUser          = "User"
	UserTypes           = []string{ "superAdmin", "Admin", "User"}
)
	

