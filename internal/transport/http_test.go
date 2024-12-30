package transport

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/poolcamacho/interviews-service/internal/domain"
	"github.com/poolcamacho/interviews-service/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Setup
	mockInterviewService := new(service.MockInterviewService)
	interviewHandler := NewInterviewHandler(mockInterviewService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/health", interviewHandler.HealthCheck)

	// Prepare HTTP request
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	// Execute
	router.ServeHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"status":"healthy"}`, rec.Body.String())
}

func TestGetInterviews(t *testing.T) {
	// Setup
	mockInterviewService := new(service.MockInterviewService)
	interviewHandler := NewInterviewHandler(mockInterviewService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/interviews", interviewHandler.GetInterviews)

	// Mock data
	interviews := []*domain.Interview{
		{
			ID:            1,
			CandidateID:   101,
			JobID:         201,
			InterviewDate: time.Date(2024, time.December, 30, 14, 0, 0, 0, time.UTC),
			Feedback:      "Excellent performance",
		},
		{
			ID:            2,
			CandidateID:   102,
			JobID:         202,
			InterviewDate: time.Date(2024, time.December, 31, 15, 0, 0, 0, time.UTC),
			Feedback:      "Needs improvement",
		},
	}

	// Mock behavior
	mockInterviewService.On("GetAllInterviews").Return(interviews, nil)

	// Prepare HTTP request
	req := httptest.NewRequest(http.MethodGet, "/interviews", nil)
	rec := httptest.NewRecorder()

	// Execute
	router.ServeHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
	responseBody, _ := json.Marshal(interviews)
	assert.JSONEq(t, string(responseBody), rec.Body.String())
	mockInterviewService.AssertExpectations(t)
}

func TestCreateInterview(t *testing.T) {
	// Setup
	mockInterviewService := new(service.MockInterviewService)
	interviewHandler := NewInterviewHandler(mockInterviewService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/interviews", interviewHandler.CreateInterview)

	// Mock data
	newInterview := &domain.Interview{
		CandidateID:   101,
		JobID:         201,
		InterviewDate: time.Date(2024, time.December, 30, 14, 0, 0, 0, time.UTC),
		Feedback:      "Good candidate",
	}

	// Mock behavior
	mockInterviewService.On("AddInterview", newInterview).Return(nil)

	// Prepare HTTP request
	body, _ := json.Marshal(newInterview)
	req := httptest.NewRequest(http.MethodPost, "/interviews", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Execute
	router.ServeHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.JSONEq(t, `{"message":"interview created successfully"}`, rec.Body.String())
	mockInterviewService.AssertExpectations(t)
}

func TestCreateInterview_BadRequest(t *testing.T) {
	// Setup
	mockInterviewService := new(service.MockInterviewService)
	interviewHandler := NewInterviewHandler(mockInterviewService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/interviews", interviewHandler.CreateInterview)

	// Invalid request body (missing required fields)
	invalidBody := `{"candidate_id": 0, "job_id": 0, "feedback": ""}`

	// Prepare HTTP request
	req := httptest.NewRequest(http.MethodPost, "/interviews", bytes.NewBufferString(invalidBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Execute
	router.ServeHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "error") // Check if the response contains an error message

	// Ensure the service method is NOT called
	mockInterviewService.AssertNotCalled(t, "AddInterview")
}
