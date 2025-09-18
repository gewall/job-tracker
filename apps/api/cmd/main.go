package main

import (
	"job-tracker/internal/handler"
	"job-tracker/internal/module"
	"github.com/gin-gonic/gin"
)

func main () {
	r := gin.Default()
	
	userModule := module.InitUser()
	
	api := r.Group("/api")
	handler.RegisterUserRoutes(api, userModule)

	
	r.Run(":8080")
}