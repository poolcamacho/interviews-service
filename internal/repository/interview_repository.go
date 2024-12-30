package repository

import (
	"database/sql"
	"github.com/poolcamacho/interviews-service/internal/domain"
)

// InterviewRepository defines methods for accessing the interviews table
// This interface abstracts database operations for the interviews table.
type InterviewRepository interface {
	// FindAll retrieves all interviews from the database
	// Executes a query to fetch all records from the interviews table.
	// @return []*domain.Interview - A slice containing all interviews
	// @return error - An error if the query fails
	FindAll() ([]*domain.Interview, error)

	// Create inserts a new interview record into the database
	// Executes an INSERT query to save a new interview in the interviews table.
	// @param interview *domain.Interview - The interview data to be saved
	// @return error - An error if the query fails
	Create(interview *domain.Interview) error
}

type interviewRepositoryImpl struct {
	db *sql.DB // Database connection instance
}

// NewInterviewRepository creates a new InterviewRepository instance
// This constructor initializes the repository with the provided database connection.
// @param db *sql.DB - The database connection used for executing queries
// @return InterviewRepository - An instance of the repository interface implementation
func NewInterviewRepository(db *sql.DB) InterviewRepository {
	return &interviewRepositoryImpl{db: db}
}

// FindAll retrieves all interviews from the database
// Executes a SELECT query on the interviews table and maps the results to a slice of Interview structs.
// @return []*domain.Interview - A slice containing all interviews
// @return error - An error if the query execution fails
func (r *interviewRepositoryImpl) FindAll() ([]*domain.Interview, error) {
	query := `SELECT id, candidate_id, job_id, interview_date, feedback FROM interviews`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err // Return error if the query fails
	}
	defer rows.Close() // Ensure rows are closed after processing

	var interviews []*domain.Interview
	for rows.Next() {
		var i domain.Interview
		// Map each row to the Interview struct
		if err := rows.Scan(&i.ID, &i.CandidateID, &i.JobID, &i.InterviewDate, &i.Feedback); err != nil {
			return nil, err // Return error if scanning fails
		}
		interviews = append(interviews, &i)
	}
	return interviews, nil
}

// Create inserts a new interview record into the database
// Executes an INSERT query to add a new record to the interviews table.
// @param interview *domain.Interview - The interview data to be saved
// @return error - An error if the query execution fails
func (r *interviewRepositoryImpl) Create(interview *domain.Interview) error {
	query := `INSERT INTO interviews (candidate_id, job_id, interview_date, feedback) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, interview.CandidateID, interview.JobID, interview.InterviewDate, interview.Feedback)
	return err // Return error if the query fails
}
