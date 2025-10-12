package service

import (
	"fmt"
	models "job-tracker/internal/model"

	"github.com/go-playground/validator/v10"
)

type JobRepositoryInterface interface {
	GetJobById(string) (*models.Application, error)
	CreateJob(*models.Application) error
	UpdateJob(*models.Application) error
	GetAllJob()(*[]models.Application,error)
}


type JobService struct {
	repo JobRepositoryInterface
}

func NewJobService(repo JobRepositoryInterface) *JobService {
	return &JobService{repo}
}

var validate = validator.New()


func (s *JobService) GetJob(id string) (*models.ApplicationResponseDTO, error) {
	job := &models.ApplicationResponseDTO{}
	res, err := s.repo.GetJobById(id)
	if err != nil {
		return nil, err
	}

	job = models.ToApplicationResponse(res)

	return job, nil
}

func (s *JobService) GetAllJob() (*[]models.ApplicationResponseDTO,error){
	datas := &[]models.ApplicationResponseDTO{}
	jobs,err := s.repo.GetAllJob()
	if err != nil {
		return nil, err
	}

	for _, job := range *jobs {
		data := models.ToApplicationResponse(&job)
		*datas = append(*datas,*data)
	}
	
	return datas, nil
}

func (s *JobService) CreateJob(data *models.ApplicationRequestDTO) error {
	job := &models.Application{}
	job = models.ToApplication(nil, data)

	if err := validate.Struct(job); err != nil {
		return err
	}
	if err := s.repo.CreateJob(job); err != nil {
		return err
	}

	return nil
}

func (s *JobService) UpdateJob(id *string, data *models.ApplicationRequestDTO) error {
	job := &models.Application{}
	if id == nil {
		return fmt.Errorf("id is invalid")
	}

	if err := validate.Struct(data); err != nil {
		return err
	}
	job = models.ToApplication(id, data)

	if err := s.repo.UpdateJob(job); err != nil {
		return err
	}

	return nil
}