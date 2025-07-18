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
	assert.Equal(t, "CREATED", *task.Status)
}

func (m *MockTaskRepository) FindAll(status string, pageSize, page int) ([]repository.Task, error) {
	args := m.Called(status, pageSize, page)
	return args.Get(0).([]repository.Task), args.Error(1)
}

func TestGetAll(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	svc := service.NewTaskService(mockRepo)
	id := uint(1)
	id2 := uint(2)
	mockTasks := []repository.Task{{ID: &id}, {ID: &id2}}
	mockRepo.On("FindAll", "MODIFIED", 10, 1).Return(mockTasks, nil)

	result, err := svc.GetAll("MODIFIED", 10, 1)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
	mockRepo.AssertExpectations(t)
}

func (m *MockTaskRepository) FindByID(id uint) (*repository.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*repository.Task), args.Error(1)
}

func TestGetByID(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	svc := service.NewTaskService(mockRepo)
	id := uint(1)
	task := &repository.Task{ID: &id}
	mockRepo.On("FindByID", uint(1)).Return(task, nil)

	result, err := svc.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), *result.ID)
	mockRepo.AssertExpectations(t)
}

func (m *MockTaskRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestDelete(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	svc := service.NewTaskService(mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := svc.Delete(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func (m *MockTaskRepository) Update(task *repository.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func TestUpdate_ValidCompletedStatus(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	svc := service.NewTaskService(mockRepo)
	id := uint(1)
	existing := &repository.Task{ID: &id}
	status := "COMPLETED"
	input := &repository.Task{ID: &id, Status: &status}

	mockRepo.On("FindByID", uint(1)).Return(existing, nil)
	mockRepo.On("Update", mock.AnythingOfType("*repository.Task")).Return(nil)

	updated, err := svc.Update(input)

	assert.NoError(t, err)
	assert.Equal(t, "COMPLETED", *updated.Status)
	mockRepo.AssertExpectations(t)
}

func TestUpdate_InvalidStatus(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	svc := service.NewTaskService(mockRepo)
	id := uint(1)
	existing := &repository.Task{ID: &id}
	status := "STARTED"
	input := &repository.Task{ID: &id, Status: &status}

	mockRepo.On("FindByID", uint(1)).Return(existing, nil)

	_, err := svc.Update(input)

	assert.Error(t, err)
	assert.Equal(t, "status should be COMPLETED", err.Error())
	mockRepo.AssertExpectations(t)
}

func ptr(s string) *string {
	return &s
}
