package repository

import (
	models "job-tracker/internal/model"

	"gorm.io/gorm"
)

type JobRepository struct {
	DB *gorm.DB
}

func NewJobRepository(DB *gorm.DB) *JobRepository {
	return &JobRepository{DB}
}

func (r *JobRepository) GetJobById(id string) (*models.Application, error){
	job := &models.Application{}
	if err := r.DB.Find(job, id); err != nil {
		return nil, err.Error
	}

	return job, nil
}

func (r *JobRepository) GetAllJob() (*[]models.Application, error){
	jobs := &[]models.Application{}
	if err := r.DB.Find(jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (r *JobRepository) CreateJob(data *models.Application) error {
	if err := r.DB.Create(data).Error; err != nil {
		return err
	}

	return nil

}

func (r *JobRepository) UpdateJob(data *models.Application) error {
	if err := r.DB.Save(data).Error; err != nil {
		return err
	}

	return nil
}