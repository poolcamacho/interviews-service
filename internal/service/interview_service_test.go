package service

import (
	"errors"
	"testing"
	"time"

	"github.com/poolcamacho/interviews-service/internal/domain"
	"github.com/poolcamacho/interviews-service/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetAllInterviews(t *testing.T) {
	// Setup
	mockRepo := new(repository.MockInterviewRepository)
	interviewService := NewInterviewService(mockRepo)

	// Mock data
	interviews := []*domain.Interview{
		{
			ID:            1,
			CandidateID:   101,
			JobID:         201,
			InterviewDate: mockInterviewDate(),
			Feedback:      "Excellent performance.",
		},
		{
			ID:            2,
			CandidateID:   102,
			JobID:         202,
			InterviewDate: mockInterviewDate(),
			Feedback:      "Needs improvement.",
		},
	}

	// Mock behavior
	mockRepo.On("FindAll").Return(interviews, nil)

	// Execute
	result, err := interviewService.GetAllInterviews()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, interviews, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAllInterviews_Error(t *testing.T) {
	// Setup
	mockRepo := new(repository.MockInterviewRepository)
	interviewService := NewInterviewService(mockRepo)

	// Mock behavior
	mockRepo.On("FindAll").Return(nil, errors.New("database error"))

	// Execute
	result, err := interviewService.GetAllInterviews()

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "database error")
	mockRepo.AssertExpectations(t)
}

func TestAddInterview(t *testing.T) {
	// Setup
	mockRepo := new(repository.MockInterviewRepository)
	interviewService := NewInterviewService(mockRepo)

	// Mock data
	newInterview := &domain.Interview{
		CandidateID:   103,
		JobID:         203,
		InterviewDate: mockInterviewDate(),
		Feedback:      "Good communication skills.",
	}

	// Mock behavior
	mockRepo.On("Create", newInterview).Return(nil)

	// Execute
	err := interviewService.AddInterview(newInterview)

	// Assertions
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddInterview_Error(t *testing.T) {
	// Setup
	mockRepo := new(repository.MockInterviewRepository)
	interviewService := NewInterviewService(mockRepo)

	// Mock data
	newInterview := &domain.Interview{
		CandidateID:   104,
		JobID:         204,
		InterviewDate: mockInterviewDate(),
		Feedback:      "Requires more technical expertise.",
	}

	// Mock behavior
	mockRepo.On("Create", newInterview).Return(errors.New("insertion error"))

	// Execute
	err := interviewService.AddInterview(newInterview)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "insertion error")
	mockRepo.AssertExpectations(t)
}

// mockInterviewDate provides a mock interview date for testing
func mockInterviewDate() time.Time {
	date, _ := time.Parse("2006-01-02 15:04:05", "2024-12-30 15:00:00")
	return date
}
