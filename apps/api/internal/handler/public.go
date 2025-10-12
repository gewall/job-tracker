package handler

import (
	"job-tracker/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PublicHandler interface {
	VerifyUser(*gin.Context) 
	
}

type publicHandler struct {
	
}

func NewPublicHandler() PublicHandler {
	return &publicHandler{}
}

func RegisterPublicHandler(r *gin.RouterGroup, h PublicHandler) {
	r.GET("/verify/:code", h.VerifyUser)
}

func (h *publicHandler) VerifyUser(c *gin.Context) {
	code := c.Param("code")

	if code == "" {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Message: "Code is invalied",
			Success: false,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Message: "Email verified!",
		Success: true,
	})
}