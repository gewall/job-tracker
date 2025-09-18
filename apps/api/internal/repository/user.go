package repository

type UserRepository interface {
	GetUser() (map[string]any, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}


func (r *userRepository) GetUser() (map[string]any, error) {
	return map[string]any{"username": "testuser"}, nil
}