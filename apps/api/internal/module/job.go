package module

import (
	"job-tracker/internal/handler"
	"job-tracker/internal/repository"
	"job-tracker/internal/service"

	"gorm.io/gorm"
)

func InitJob(DB *gorm.DB) *handler.JobHandler {
	repo := repository.NewJobRepository(DB)
	svc := service.NewJobService(repo)
	return handler.NewJobHandler(svc)
}