package service

import (
	"fmt"
	"job-tracker/internal/config"
	models "job-tracker/internal/model"
	"job-tracker/internal/repository"
	"job-tracker/internal/seeder"
	"job-tracker/internal/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUser() (*models.UserResponseDTO, error)
	SeedUser() error
	VerifyUser(code string) (*models.UserResponseDTO, error)
	Signin(*models.UserSigninDTO) (*models.TokenResponseDTO,error)
}

type userService struct {
	repository repository.UserRepository
	email *utils.Email
	config *config.Config
}

func NewUserService(repository repository.UserRepository, email *utils.Email, cofing *config.Config) UserService {
	return &userService{
		repository: repository,
		email: email,
		config: cofing,
	}
}

func (s *userService) Signin(user *models.UserSigninDTO) (*models.TokenResponseDTO, error) {
	res, err := s.repository.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	data := &models.UserSigninDTO{
		Email: res.Email,
		Password: res.Password,
	}
	
	
	
	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(user.Password)); err != nil {
		return nil, err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"userId": res.ID,
		"email": data.Email,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	}

	tokenString, err := token.SignedString([]byte(s.config.JwtSecretKey))

	if err != nil {
		return nil, err
	}

	refreshToken := utils.GenerateRefreshToken(int(res.ID))
	refreshTokenData := &models.RefreshToken{
		UserID: res.ID,
		Token: refreshToken.Token,
		ExpiresAt: refreshToken.ExpiresAt,
	}

	if err := s.repository.CreateRefreshToken(refreshTokenData); err != nil {
		return nil, err
	}

	

	return &models.TokenResponseDTO{
		RefreshToken: refreshToken.Token,
		ExpiresAt: refreshToken.ExpiresAt,
		AccessToken: tokenString,
	},nil
}

func (s *userService) SeedUser() error {
	seeder := seeder.New(s.repository)
	res,err := seeder.SeedUser(); 
	if err != nil {
		return err
	}

	go s.email.SendEmail(utils.Mail{
		From: "alginugraha21@gmail.com",
		To: res.Email,
		Subject: "Job Tracker",
		Body: fmt.Sprintf("<h4>Email Verification</h4><br></br><p>Use this link for your verification, <a href='http://localhost:8080/verify/%s'>Click this link!</a></p>", res.Verifications[0].Code,
		),
		// Body: "Your email is registered as a user",
	})

	return nil
}

func (s *userService) GetUser() (*models.UserResponseDTO, error) {
	res, err := s.repository.GetUser("1")
	user := &models.UserResponseDTO{
		Name: res.Name,
		ID: res.ID,
		Email: res.Email,
		IsVerified: res.IsVerified,
		Verifications: []models.VerifyEmailDTO{
			{
				ExpiresAt: res.Verifications[0].ExpiresAt,
				IsUsed: res.Verifications[0].IsUsed,
				Code: res.Verifications[0].Code,
			},
		},
	}
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (s *userService) VerifyUser(code string) (*models.UserResponseDTO, error) {
	user := &models.UserResponseDTO{}
	res, err := s.repository.GetUserEmailVerification(code)

	if err != nil {
		return nil, err
	}

	user = &models.UserResponseDTO{
		ID: res.ID,
		Name: res.Name,
		Email: res.Email,
		IsVerified: res.IsVerified,
		Verifications: []models.VerifyEmailDTO{
			{
				ExpiresAt: res.Verifications[0].ExpiresAt,
				Code: res.Verifications[0].Code,
				IsUsed: res.Verifications[0].IsUsed,
			},
		},
	}



	return user, nil
}

func (s *userService) SendEmailVerification() {
	
}