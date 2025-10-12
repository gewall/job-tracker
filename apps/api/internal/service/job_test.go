package service_test

import (
	"errors"
	models "job-tracker/internal/model"
	"job-tracker/internal/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockJobRepository struct {}

func (m *mockJobRepository) GetJobById(id string) (*models.Application, error) {
	if id == "1" {
		return &models.Application{
			ID:              "a3b2f4c6-9d2e-4f8e-b123-8e71b29f4567",
			CompanyName:     "TechNova Solutions",
			Position:        "Backend Developer",
			ApplicationDate: time.Date(2025, 10, 1, 0, 0, 0, 0, time.UTC),
			Status:          "interview",
			InterviewDate:   ptrTime(time.Date(2025, 10, 15, 10, 0, 0, 0, time.UTC)),
			InterviewLocation: ptrString("Zoom Meeting - Link sent via email"),
			ContactHR:         ptrString("hr@technova.com"),
			SalaryExpectation: ptrFloat64(12000000),
			SalaryOffered:     ptrFloat64(12500000),
			Location:          ptrString("Jakarta"),
			JobType:           ptrString("full-time"),
			Feedback:          ptrString("Interview went well, waiting for final decision"),
			Notes:             ptrString("Company culture seems good, salary slightly above expectation"),
			ResumeVersion:     ptrString("v3.1"),

			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	}

	return &models.Application{}, errors.New("Job not found")
}
func (m *mockJobRepository) CreateJob(data *models.ApplicationRequestDTO)error
func (m *mockJobRepository) UpdateJob(data *models.Application)error
func (m *mockJobRepository) GetAllJob()(*[]models.Application,error)

func ptrString(s string) *string       { return &s }
func ptrFloat64(f float64) *float64    { return &f }
func ptrTime(t time.Time) *time.Time   { return &t }


func TestGetUser(t *testing.T) {
	mockRepo := &mockJobRepository{}
	service := service.NewJobService(mockRepo)

	user, err := service.GetJob("1")
	assert.NoError(t, err)
	assert.Equal(t,"TechNova Solutions", user.CompanyName)

	_, err = service.GetJob("123")
	assert.Error(t, err)
	
}