package routers

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/utils/encrypts"
	"PetPalApp/utils/helper"

	_adminData "PetPalApp/features/admin/data"
	_adminHandler "PetPalApp/features/admin/handler"
	_adminService "PetPalApp/features/admin/service"
	_userData "PetPalApp/features/user/data"
	_userHandler "PetPalApp/features/user/handler"
	_userService "PetPalApp/features/user/service"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB, s3 *s3.S3, s3Bucket string) {
	hashService := encrypts.NewHashService()
	helperService := helper.NewHelperService(s3, s3Bucket)

	userData := _userData.New(db, helperService)
	userService := _userService.New(userData, hashService, helperService)
	userHandlerAPI := _userHandler.New(userService, hashService)

	adminData := _adminData.New(db)
	adminService := _adminService.New(adminData, hashService)
	adminHandlerAPI := _adminHandler.New(adminService)


	//user
	e.POST("/user/register", userHandlerAPI.Register)
	e.POST("/user/login", userHandlerAPI.Login)
	e.GET("/user/profile", userHandlerAPI.Profile, middlewares.JWTMiddleware())
	e.PUT("/user", userHandlerAPI.UpdateUserById, middlewares.JWTMiddleware())

	//admin
	e.POST("/admin/register", adminHandlerAPI.Register)
}