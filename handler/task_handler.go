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

func (h *TaskHandler) GetAll(c *gin.Context) {
	status := c.Query("status")
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page number must be greater than 0"})
		return
	}
	if pageSize < 0 || pageSize > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page size must be between 1 and 100 "})
		return
	}
	tasks, err := h.service.GetAll(status, pageSize, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (r *taskRepository) FindByID(id uint) (*Task, error) {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
