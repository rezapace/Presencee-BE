package bootsrapper

import (
	"time"

	userControllerPkg "presensee_project/controller"
	userRepositoryPkg "presensee_project/repository/impl"
	routes "presensee_project/route"
	userServicePkg "presensee_project/usecase/impl"
	jwtPkg "presensee_project/utils/jwt_service/impl"
	passwordPkg "presensee_project/utils/password/impl"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

func InitController(e *echo.Echo, db *gorm.DB, conf map[string]string) {
	passwordFunc := passwordPkg.NewPasswordFuncImpl()
	jwtService := jwtPkg.NewJWTService(conf["JWT_SECRET"], 1*time.Hour)

	// User
	userRepository := userRepositoryPkg.NewUserRepositoryImpl(db)
	userService := userServicePkg.NewUserServiceImpl(userRepository, passwordFunc, jwtService)
	userController := userControllerPkg.NewUserController(userService, jwtService)

	route := routes.NewRoutes(userController)
	route.Init(e, conf)
}
