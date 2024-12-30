package service

import (
	"github.com/poolcamacho/interviews-service/internal/domain"
	"github.com/stretchr/testify/mock"
)

// MockInterviewService is a mock implementation of InterviewService for testing
type MockInterviewService struct {
	mock.Mock
}

// GetAllInterviews mocks the GetAllInterviews method
// @return []*domain.Interview - A slice of interviews
// @return error - An error if the operation fails
func (m *MockInterviewService) GetAllInterviews() ([]*domain.Interview, error) {
	args := m.Called()
	if interviews, ok := args.Get(0).([]*domain.Interview); ok {
		return interviews, args.Error(1)
	}
	return nil, args.Error(1)
}

// AddInterview mocks the AddInterview method
// @param interview *domain.Interview - The interview data to be added
// @return error - An error if the operation fails
func (m *MockInterviewService) AddInterview(interview *domain.Interview) error {
	args := m.Called(interview)
	return args.Error(0)
}

// GetInterviewByID mocks the GetInterviewByID method
// @param id int - The ID of the interview to retrieve
// @return *domain.Interview - The retrieved interview
// @return error - An error if the operation fails
func (m *MockInterviewService) GetInterviewByID(id int) (*domain.Interview, error) {
	args := m.Called(id)
	if interview, ok := args.Get(0).(*domain.Interview); ok {
		return interview, args.Error(1)
	}
	return nil, args.Error(1)
}

// UpdateInterview mocks the UpdateInterview method
// @param interview *domain.Interview - The interview data to be updated
// @return error - An error if the operation fails
func (m *MockInterviewService) UpdateInterview(interview *domain.Interview) error {
	args := m.Called(interview)
	return args.Error(0)
}

// DeleteInterview mocks the DeleteInterview method
// @param id int - The ID of the interview to delete
// @return error - An error if the operation fails
func (m *MockInterviewService) DeleteInterview(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
