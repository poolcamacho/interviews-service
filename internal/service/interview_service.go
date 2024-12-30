package service

import (
	"github.com/poolcamacho/interviews-service/internal/domain"
	"github.com/poolcamacho/interviews-service/internal/repository"
)

// InterviewService defines methods for interview-related operations
// This interface abstracts the business logic for managing interviews.
type InterviewService interface {
	// GetAllInterviews retrieves all interviews from the repository
	// Delegates the retrieval operation to the repository layer.
	// @return []*domain.Interview - A slice containing all interviews
	// @return error - An error if there is an issue retrieving the interviews
	GetAllInterviews() ([]*domain.Interview, error)

	// AddInterview adds a new interview to the repository
	// Delegates the creation operation to the repository layer.
	// @param interview *domain.Interview - The interview data to be added
	// @return error - An error if there is an issue creating the interview
	AddInterview(interview *domain.Interview) error
}

type interviewServiceImpl struct {
	repo repository.InterviewRepository // Dependency on the InterviewRepository
}

// NewInterviewService creates a new InterviewService instance
// This constructor initializes the service with the provided repository.
// @param repo repository.InterviewRepository - The repository used for database operations
// @return InterviewService - An instance of the service interface implementation
func NewInterviewService(repo repository.InterviewRepository) InterviewService {
	return &interviewServiceImpl{repo: repo}
}

// GetAllInterviews retrieves all interviews from the repository
// This method interacts with the repository layer to fetch all interview records.
// @return []*domain.Interview - A slice containing all interviews
// @return error - An error if the retrieval operation fails
func (s *interviewServiceImpl) GetAllInterviews() ([]*domain.Interview, error) {
	return s.repo.FindAll() // Call the repository method to fetch all interviews
}

// AddInterview adds a new interview to the repository
// This method interacts with the repository layer to save a new interview record.
// @param interview *domain.Interview - The interview data to be added
// @return error - An error if the creation operation fails
func (s *interviewServiceImpl) AddInterview(interview *domain.Interview) error {
	return s.repo.Create(interview) // Call the repository method to add the new interview
}
