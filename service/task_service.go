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
	status := string(repository.StatusCreated)
	task.Status = &status	
	return s.repo.Create(task)
}

func (s *taskService) GetAll(status string, pageSize, page int) ([]repository.Task, error) {
	return s.repo.FindAll(status, pageSize, page)
}

func (s *taskService) GetByID(id uint) (*repository.Task, error) {
	return s.repo.FindByID(id)
}

func (s *taskService) Delete(id uint) error {
	return s.repo.Delete(id)
}


func (s *taskService) Update(input *repository.Task) (*repository.Task, error) {
	existingTask, err := s.repo.FindByID(*input.ID)
	if err != nil {
		return nil, fmt.Errorf("task not found")
	}
	if input.Title != nil {
		existingTask.Title = input.Title
	}
	if input.Description != nil {
		existingTask.Description = input.Description
	}
	if input.Status != nil {
		if strings.ToUpper(*input.Status) != string(repository.StatusCompleted) {
			return nil, fmt.Errorf("status should be COMPLETED")
		}
		existingTask.Status = input.Status
	} else {
		status := string(repository.StatusModified)
		existingTask.Status = &status
	}
	if err := s.repo.Update(existingTask); err != nil {
		return nil, err
	}
	return existingTask, nil
}