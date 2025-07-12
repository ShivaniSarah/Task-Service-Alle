package service_test

import (
	"testing"
	"task-service/repository"
	"task-service/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Mock Repository ---
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Create(task *repository.Task) error {
	args := m.Called(task)
	return args.Error(0)
}




func TestCreate(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	svc := service.NewTaskService(mockRepo)

	task := &repository.Task{Title: ptr("Test")}
	mockRepo.On("Create", mock.AnythingOfType("*repository.Task")).Return(nil)

	err := svc.Create(task)

	assert.NoError(t, err)
	assert.Equal(t, "MODIFIED", *task.Status)
	mockRepo.AssertExpectations(t)
}



func ptr(s string) *string {
	return &s
}
