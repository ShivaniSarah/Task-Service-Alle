package repository

import "gorm.io/gorm"

type TaskRepository interface {
	Create(task *Task) error
	FindAll(status string, limit, offset int) ([]Task, error)
	FindByID(id uint) (*Task, error)
	Update(task *Task) error
	Delete(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *Task) error {
	return r.db.Create(task).Error
}

