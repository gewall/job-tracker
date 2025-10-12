package module

import (
	"job-tracker/internal/config"
	"job-tracker/internal/handler"
	"job-tracker/internal/repository"
	"job-tracker/internal/service"
	"job-tracker/internal/utils"

	"gorm.io/gorm"
)


func InitUser(email *utils.Email, config *config.Config, db *gorm.DB) handler.UserHandler {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, email,config)

	return handler.NewUserHandler(userService,email,config)
}