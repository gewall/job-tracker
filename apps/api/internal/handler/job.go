package handler

import (
	"fmt"
	models "job-tracker/internal/model"
	"job-tracker/internal/service"
	"job-tracker/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	service *service.JobService
}

func NewJobHandler(service *service.JobService) *JobHandler {
	return &JobHandler{service}
}

func RegisterJobHandler(r *gin.RouterGroup, h *JobHandler) {
	r.GET("job/:id", h.GetJob)
	r.POST("/job", h.CreateJob)
	r.GET("/job", h.GetAllJob)
	r.PATCH("/job/:id", h.UpdateJob)
}

func (h *JobHandler) GetJob(c *gin.Context) {
	id := c.Param("id")

	if id != "" {
		c.JSON(http.StatusBadRequest, utils.Response{
			Message: "id is invalid",
			Success: false,
		})
		return
	}

	res, err := h.service.GetJob(id)
	if err != nil {
			c.JSON(http.StatusInternalServerError, utils.Response{
			Message: fmt.Sprintf("fail to fetch job: %s", err),
			Success: false,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Message: "fetch successfully",
		Data: res,
		Success: true,
	})

}

func (h *JobHandler) GetAllJob(c *gin.Context) {
	data, err := h.service.GetAllJob()
	if err != nil {
			c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Message: "request successfully",
		Data: data,
		Success: true,
	})
}

func (h *JobHandler) UpdateJob(c *gin.Context) {
	id := c.Param("id")
	data := models.ApplicationRequestDTO{}
	if id == "" {
		c.JSON(http.StatusBadRequest, utils.Response{
			Message: "id invalid",
			Success: false,
		})
		return
	}

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	if err := h.service.UpdateJob(&id, &data); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Message: "request successfully",
		Success: true,
		
	})
	
}

func (h *JobHandler) CreateJob(c *gin.Context) {
	var data models.ApplicationRequestDTO

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Message: "request invalid",
			Success: false,
		})
		return
	}

	err := h.service.CreateJob(&data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Message: fmt.Sprintf("fail to create job: %s", err),
			Success: false,
		})
		return
	}
	
	
}