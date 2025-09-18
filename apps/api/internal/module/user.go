package module

import (
	"job-tracker/internal/handler"
	"job-tracker/internal/repository"
	"job-tracker/internal/service"
)


func InitUser() handler.UserHandler {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)

	return handler.NewUserHandler(userService)
}