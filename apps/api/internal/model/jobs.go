package models

import (
	"time"

	"gorm.io/gorm"
)

type ApplicationStatus string

const (
	StatusApplied    ApplicationStatus = "applied"
	StatusInterview  ApplicationStatus = "interview"
	StatusOffered    ApplicationStatus = "offered"
	StatusRejected   ApplicationStatus = "rejected"
	StatusAccepted   ApplicationStatus = "accepted"
)

type Application struct {
	ID                string            `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CompanyName       string            `gorm:"not null" json:"company_name" validate:"required"`
	Position          string            `gorm:"not null" json:"position" validate:"required"`
	ApplicationDate   time.Time         `gorm:"not null" json:"application_date" validate:"required"`
	Status            ApplicationStatus `gorm:"type:varchar(20);not null" json:"status" validate:"required,oneof=applied interview offered rejected accepted"`
	InterviewDate     *time.Time        `json:"interview_date"`
	InterviewLocation *string           `json:"interview_location"`
	ContactHR         *string           `json:"contact_hr"`
	SalaryExpectation *float64          `json:"salary_expectation" validate:"omitempty,gte=0"`
	SalaryOffered     *float64          `json:"salary_offered" validate:"omitempty,gte=0"`
	Location          *string           `json:"location"`
	JobType           *string           `gorm:"type:varchar(20)" json:"job_type" validate:"omitempty,oneof=full-time part-time internship contract"`
	Feedback          *string           `json:"feedback"`
	Notes             *string           `json:"notes"`
	ResumeVersion     *string           `json:"resume_version"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}


type ApplicationRequestDTO struct {
	CompanyName       string     `json:"company_name" validate:"required"`
	Position          string     `json:"position" validate:"required"`
	ApplicationDate   time.Time  `json:"application_date" validate:"required"`
	Status            string     `json:"status" validate:"required,oneof=applied interview offered rejected accepted"`
	InterviewDate     *time.Time `json:"interview_date,omitempty"`
	InterviewLocation *string    `json:"interview_location,omitempty"`
	ContactHR         *string    `json:"contact_hr,omitempty"`
	SalaryExpectation *float64   `json:"salary_expectation,omitempty" validate:"omitempty,gte=0"`
	SalaryOffered     *float64   `json:"salary_offered,omitempty" validate:"omitempty,gte=0"`
	Location          *string    `json:"location,omitempty"`
	JobType           *string    `json:"job_type,omitempty" validate:"omitempty,oneof=full-time part-time internship contract"`
	Feedback          *string    `json:"feedback,omitempty"`
	Notes             *string    `json:"notes,omitempty"`
	ResumeVersion     *string    `json:"resume_version,omitempty"`
}


type ApplicationResponseDTO struct {
	ID                string     `json:"id"`
	CompanyName       string     `json:"company_name"`
	Position          string     `json:"position"`
	ApplicationDate   time.Time  `json:"application_date"`
	Status            string     `json:"status"`
	InterviewDate     *time.Time `json:"interview_date,omitempty"`
	InterviewLocation *string    `json:"interview_location,omitempty"`
	ContactHR         *string    `json:"contact_hr,omitempty"`
	SalaryExpectation *float64   `json:"salary_expectation,omitempty"`
	SalaryOffered     *float64   `json:"salary_offered,omitempty"`
	Location          *string    `json:"location,omitempty"`
	JobType           *string    `json:"job_type,omitempty"`
	Feedback          *string    `json:"feedback,omitempty"`
	Notes             *string    `json:"notes,omitempty"`
	ResumeVersion     *string    `json:"resume_version,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

func ToApplication(id *string, app *ApplicationRequestDTO) *Application {
	if id != nil {
		return &Application{
			ID:                ToStringPTR(id),
			CompanyName:       app.CompanyName,
			Position:          app.Position,
			ApplicationDate:   app.ApplicationDate,
			Status:            ApplicationStatus(app.Status),
			InterviewDate:     app.InterviewDate,
			InterviewLocation: app.InterviewLocation,
			ContactHR:         app.ContactHR,
			SalaryExpectation: app.SalaryExpectation,
			SalaryOffered:     app.SalaryOffered,
			Location:          app.Location,
			JobType:           app.JobType,
			Feedback:          app.Feedback,
			Notes:             app.Notes,
			ResumeVersion:     app.ResumeVersion,
		}
		}	
	return &Application{
		
		CompanyName:       app.CompanyName,
		Position:          app.Position,
		ApplicationDate:   app.ApplicationDate,
		Status:            ApplicationStatus(app.Status),
		InterviewDate:     app.InterviewDate,
		InterviewLocation: app.InterviewLocation,
		ContactHR:         app.ContactHR,
		SalaryExpectation: app.SalaryExpectation,
		SalaryOffered:     app.SalaryOffered,
		Location:          app.Location,
		JobType:           app.JobType,
		Feedback:          app.Feedback,
		Notes:             app.Notes,
		ResumeVersion:     app.ResumeVersion,
	}
}

func ToStringPTR(a *string) string {
	return *a
}

func ToApplicationResponse(app *Application) *ApplicationResponseDTO {
	return &ApplicationResponseDTO{
		ID:                app.ID,
		CompanyName:       app.CompanyName,
		Position:          app.Position,
		ApplicationDate:   app.ApplicationDate,
		Status:            string(app.Status),
		InterviewDate:     app.InterviewDate,
		InterviewLocation: app.InterviewLocation,
		ContactHR:         app.ContactHR,
		SalaryExpectation: app.SalaryExpectation,
		SalaryOffered:     app.SalaryOffered,
		Location:          app.Location,
		JobType:           app.JobType,
		Feedback:          app.Feedback,
		Notes:             app.Notes,
		ResumeVersion:     app.ResumeVersion,
		CreatedAt:         app.CreatedAt,
		UpdatedAt:         app.UpdatedAt,
	}
}