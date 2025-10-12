package handler

import (
	"fmt"
	"job-tracker/internal/config"
	models "job-tracker/internal/model"
	"job-tracker/internal/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Auth interface {
	SignIn(c *gin.Context)
}

type auth struct {
 config *config.Config
}

func NewAuth(cfg *config.Config) Auth {
	return &auth{
		config: cfg,
	}
}

func RegisterAuth(r *gin.RouterGroup, h Auth) {
	r.POST("/signin", h.SignIn)
}

func (a auth) SignIn(c *gin.Context) {
	data := models.UserResponseDTO{}

	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Message: fmt.Sprintf("fail to bind request: %s", err.Error()),
			Success: false,
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"userId": data.ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	tokenString,err := token.SignedString([]byte(a.config.JwtSecretKey))
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Message: fmt.Sprintf("fail to generate token: %s", err.Error()),
			Success: false,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Message: "Signin Successfully",
		Success: true,
		Data: map[string]string{
			"token": tokenString,
		},
	})
}