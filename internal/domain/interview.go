package domain

import "time"

// Interview represents an interview record in the system
// This struct defines the schema of an interview as it is stored in the database.
type Interview struct {
	ID            int       `json:"id"`             // Unique identifier for the interview
	CandidateID   int       `json:"candidate_id"`   // Foreign key referencing the candidate's ID
	JobID         int       `json:"job_id"`         // Foreign key referencing the job's ID
	InterviewDate time.Time `json:"interview_date"` // Date and time of the interview
	Feedback      string    `json:"feedback"`       // Feedback or notes about the interview
}
