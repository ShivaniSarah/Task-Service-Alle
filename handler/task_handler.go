package handler

import (
	"net/http"
	"strconv"
	"task-service/repository"
	"task-service/service"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(s service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/tasks", h.GetAll)
	r.POST("/tasks", h.Create)
	r.GET("/tasks/:id", h.GetByID)
	r.PATCH("/task", h.Update)
	r.DELETE("/tasks/:id", h.Delete)
}

func (h *TaskHandler) Create(c *gin.Context) {
	var task repository.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.service.Create(&task)
	c.JSON(http.StatusCreated, task)
}


