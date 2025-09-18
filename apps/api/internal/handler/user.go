package handler

import (
	"job-tracker/internal/service"
	"job-tracker/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	PingUser(c *gin.Context)
	GetUser(c *gin.Context)
}

type userHandler struct{
	Service service.UserService
}

func NewUserHandler(s service.UserService) UserHandler {
	return &userHandler{
		Service: s,
	}
}

func RegisterUserRoutes(r *gin.RouterGroup, h UserHandler) {
	r.GET("/pinguser", h.PingUser)
}

func (h *userHandler) PingUser(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Response{
		Message: "User Pong",
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
		Message: "User fetched successfully",
		Data:    user,
		Success: true,
	})
}