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

func (r *taskRepository) FindAll(status string, limit, offset int) ([]Task, error) {
	var tasks []Task
	realOffset := (offset - 1) * limit
	query := r.db.Limit(limit).Offset(realOffset)
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) FindByID(id uint) (*Task, error) {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}