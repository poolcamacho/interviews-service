package repository

import (
	"github.com/poolcamacho/interviews-service/internal/domain"
	"github.com/stretchr/testify/mock"
)

// MockInterviewRepository is a mock implementation of InterviewRepository for testing
type MockInterviewRepository struct {
	mock.Mock
}

// FindAll mocks the FindAll method
// @return []*domain.Interview - A slice of interviews
// @return error - An error if the operation fails
func (m *MockInterviewRepository) FindAll() ([]*domain.Interview, error) {
	args := m.Called()
	if interviews, ok := args.Get(0).([]*domain.Interview); ok {
		return interviews, args.Error(1)
	}
	return nil, args.Error(1)
}

// Create mocks the Create method
// @param interview *domain.Interview - The interview data to be added
// @return error - An error if the operation fails
func (m *MockInterviewRepository) Create(interview *domain.Interview) error {
	args := m.Called(interview)
	return args.Error(0)
}
