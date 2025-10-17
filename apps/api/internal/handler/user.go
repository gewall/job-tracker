package handler

import (
	"fmt"
	"job-tracker/internal/config"
	models "job-tracker/internal/model"
	"job-tracker/internal/service"
	"job-tracker/internal/utils"
	"log"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	PingUser(c *gin.Context)
	GetUser(c *gin.Context)
	SeedUser(c *gin.Context)
	VerifyUser(c *gin.Context)
	Signin(c *gin.Context)
}

type userHandler struct{
	Service service.UserService
	Email *utils.Email
	Config *config.Config
}

func NewUserHandler(service service.UserService,email *utils.Email,config *config.Config,  ) UserHandler {
	return &userHandler{
		Service: service,
		Email: email,
		Config: config,
	
	}
}

func RegisterUserRoutes(r *gin.RouterGroup,p *gin.RouterGroup, h UserHandler) {
	r.GET("/pinguser", h.PingUser)
	r.POST("/user/seed", h.SeedUser)
	r.GET("/user", h.GetUser)
	p.GET("/verify/:code", h.VerifyUser)
	p.POST("/auth/signin", h.Signin)
}

func (h *userHandler) Signin(c *gin.Context) {
	user := &models.UserSigninDTO{}

	if err := c.BindJSON(user); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	log.Println("em", user.Email)

	res,err := h.Service.Signin(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Success: false,
		})
		return
	}
	
	c.SetCookie(
    "refresh_token",
    res.RefreshToken,
    int((time.Hour * 24 * 7).Seconds()), // 7 hari
    "/",
    "",      // domain kosong
    false,   // secure = false (karena localhost)
    true,    // httpOnly = true
)

	c.JSON(http.StatusOK, utils.Response{
		Message: "Signin successfully",
		Success: true,
		Data: map[string]string{
			"accessToken": res.AccessToken, 
		},
	})

}

func (h *userHandler) PingUser(c *gin.Context) {

	go h.Email.SendEmail(utils.Mail{
		From: h.Config.Email.Email,
		To: "gewall21@gmail.com",
		Subject: "Test Golang SMTP",
		Body: "<h2>Percobaan mengirim email dengan golang<h2>",
	})

	// go h.Email.SendEmail(h.Config.Email.Email,"gewall21@gmail.com","TES GOLANG","Adalah saya")


	c.JSON(http.StatusOK, utils.Response{
		Message: "User Pong",
		Success: true,
	})
}

func (h *userHandler) SeedUser(c *gin.Context) {
	if err := h.Service.SeedUser(); err != nil {
		c.JSON(http.StatusInternalServerError,utils.Response{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	c.JSON(http.StatusCreated,utils.Response{
		Message: "Seed user successfully",
		Success: true,
	})
}

func (h *userHandler) GetUser(c *gin.Context) {

	user, err := h.Service.GetUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Success: false,
		})
		return
	}
	

	c.JSON(http.StatusOK, utils.Response{
		Message: "User fetched successfully oke",
		Data:    user,
		Success: true,
	})
}


func (h *userHandler) VerifyUser(c *gin.Context) {
	code := c.Param("code")

	if code == "" {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Message: "Code is invalied",
			Success: false,
		})
		return
	}

	res, err := h.Service.VerifyUser(code)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Message: fmt.Sprintf("Fail to verify code: %s", err),
			Success: false,
		})
	}

	c.JSON(http.StatusOK, utils.Response{
		Message: fmt.Sprintf("Email %s verified!", res.Email),
		Success: true,
	})
}

func (h *userHandler) SendEmailVerification(c *gin.Context) {

}