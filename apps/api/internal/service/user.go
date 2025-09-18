package service

import "job-tracker/internal/repository"

type UserService interface {
	GetUser() (map[string]any, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		repository: r,
	}
}

func (s *userService) GetUser() (map[string]any, error) {
	return s.repository.GetUser()
}