package service

import (
	"fmt"
	"strings"
	"task-service/repository"
)

type TaskService interface {
	Create(task *repository.Task) error
	GetAll(status string, pageSize, page int) ([]repository.Task, error)
	GetByID(id uint) (*repository.Task, error)
	Update(task *repository.Task) (*repository.Task, error)
	Delete(id uint) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) Create(task *repository.Task) error {
	status := string(repository.StatusModified)
	task.Status = &status
	return s.repo.Create(task)
}

func (s *taskService) GetAll(status string, pageSize, page int) ([]repository.Task, error) {
	return s.repo.FindAll(status, pageSize, page)
}

func (s *taskService) GetByID(id uint) (*repository.Task, error) {
	return s.repo.FindByID(id)
}