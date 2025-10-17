package repository

import (
	"errors"
	"fmt"
	models "job-tracker/internal/model"

	"time"

	"strconv"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(ID string) (*models.User, error)
	CreateUser(user *models.User) (string,error)
	GetUserEmailVerification(string) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	CreateRefreshToken(*models.RefreshToken) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error){
	user := models.User{
	
	}

	res := r.DB.Limit(1).Where("Email = ?", email).First(&user);

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (r *userRepository) CreateRefreshToken(token *models.RefreshToken) error {
	data := &models.RefreshToken{
		UserID: token.UserID,
		Token: token.Token,
		ExpiresAt: token.ExpiresAt,
	}

	if err := r.DB.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository)CreateUser(user *models.User) (string, error) {
	usr := user
	res := r.DB.Create(&usr)

	
	if res.Error != nil {
		return "", fmt.Errorf("fail to create user: %s", res.Error)
	}

	return  strconv.FormatUint(uint64(usr.ID),10),nil
}

func (r *userRepository) GetUser(ID string) (*models.User, error) {
	user := &models.User{}
	err := r.DB.Preload("Verifications").Find(user,ID).Error
	fmt.Println(user.Verifications)

	if err!= nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserEmailVerification(code string) (*models.User, error) {
	user := &models.User{} 
	if err := r.DB.Preload("Verifications","Code = ? AND Expires_At > ? AND Is_Used = ?", code, time.Now(), false).Find(user).Error; err != nil {
		return nil, err
	}

	

	return user, nil
}

func (r *userRepository) SendEmailVerification(ID interface{}) {
	
}