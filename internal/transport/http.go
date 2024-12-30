package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/poolcamacho/interviews-service/internal/domain"
	"github.com/poolcamacho/interviews-service/internal/service"
)

// InterviewHandler handles HTTP requests for interviews
type InterviewHandler struct {
	service service.InterviewService
}

// NewInterviewHandler creates a new InterviewHandler instance
// @Summary Initialize the interview handler
// @Description Creates an instance of InterviewHandler to manage interview endpoints
// @Tags Initialization
// @Produce json
func NewInterviewHandler(service service.InterviewService) *InterviewHandler {
	return &InterviewHandler{service: service}
}

// HealthCheck provides a simple health status of the service
// @Summary Check service health
// @Description Returns the health status of the interview service
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string "Service is healthy"
// @Router /health [get]
func (h *InterviewHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

// GetInterviews handles the retrieval of all interviews
// @Summary Get all interviews
// @Description Retrieve a list of all interviews in the system
// @Tags Interviews
// @Produce json
// @Success 200 {array} domain.Interview "List of interviews"
// @Failure 500 {object} map[string]string "Failed to fetch interviews"
// @Router /interviews [get]
func (h *InterviewHandler) GetInterviews(c *gin.Context) {
	interviews, err := h.service.GetAllInterviews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch interviews"})
		return
	}
	c.JSON(http.StatusOK, interviews)
}

// CreateInterview handles the creation of a new interview
// @Summary Create a new interview
// @Description Add a new interview by providing candidate_id, job_id, interview_date, and feedback
// @Tags Interviews
// @Accept json
// @Produce json
// @Param request body domain.Interview true "Interview Creation Request"
// @Success 201 {object} map[string]string "Interview created successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Failed to create interview"
// @Router /interviews [post]
func (h *InterviewHandler) CreateInterview(c *gin.Context) {
	var interview domain.Interview
	if err := c.ShouldBindJSON(&interview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields
	if interview.CandidateID == 0 || interview.JobID == 0 || interview.InterviewDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "candidate_id, job_id, and interview_date are required"})
		return
	}

	// Call the service to add the interview
	if err := h.service.AddInterview(&interview); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create interview"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "interview created successfully"})
}
