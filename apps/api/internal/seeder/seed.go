package seeder

import (
	"fmt"
	models "job-tracker/internal/model"
	"job-tracker/internal/repository"
	"job-tracker/internal/utils"

	"golang.org/x/crypto/bcrypt"
)
type Seeder struct {
	repository repository.UserRepository
}

func New(repository repository.UserRepository) *Seeder{
	return&Seeder{
		repository: repository,
	}
}

func (s *Seeder) SeedUser() (*models.User,error) {
	// usr := &models.User{}
	verifCode,_ := utils.GenerateVerificationCode()
	user := &models.User{
		Name: "refa", Email: "rafderilera@gmail.com", IsVerified: false,
		Verifications: []models.EmailVerification{
			{Code: verifCode.Token,
			ExpiresAt:  verifCode.ExpiresAt},
		},
	}
	pass := "password123"

	hash,err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if err != nil {
		return nil,fmt.Errorf("fail to hash password")
	}

	user.Password = string(hash)

	id,err := s.repository.CreateUser(user);
	if err != nil {
		return nil,err
	}

	res,err :=  s.repository.GetUser(id)
	if err != nil {
		return nil, err
	}


	return res,nil
}