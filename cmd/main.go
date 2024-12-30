package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/poolcamacho/interviews-service/internal/repository"
	"github.com/poolcamacho/interviews-service/internal/service"
	"github.com/poolcamacho/interviews-service/internal/transport"
	"github.com/poolcamacho/interviews-service/pkg/config"
	"github.com/poolcamacho/interviews-service/pkg/db"
	jwtUtil "github.com/poolcamacho/interviews-service/pkg/jwt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/poolcamacho/interviews-service/docs" // Import Swagger docs
)

// @title Interview Service API
// @version 1.0
// @description API for managing interviews in the system.
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

// main initializes and starts the Interview Service API server
// @Summary Start the Interview Service
// @Description Initializes the Interview Service API with routes and Swagger documentation.
// @Tags Initialization
func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to the database
	dbConn := db.Connect(cfg.DatabaseURL)

	// Initialize repository
	interviewRepository := repository.NewInterviewRepository(dbConn)

	// Initialize service
	interviewService := service.NewInterviewService(interviewRepository)

	// Initialize Gin and routes
	r := gin.Default()
	handler := transport.NewInterviewHandler(interviewService)

	// Swagger route
	// @Summary Swagger Documentation
	// @Description Provides Swagger UI for the API
	// @Tags Swagger
	// @Router /swagger/*any [get]
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register routes
	// @Summary Get all interviews
	// @Description Fetch a list of all interviews
	// @Tags Interviews
	// @Produce json
	// @Success 200 {array} domain.Interview
	// @Router /interviews [get]
	r.GET("/interviews", jwtUtil.AuthMiddleware(cfg.JWTSecretKey), handler.GetInterviews)

	// @Summary Create a new interview
	// @Description Add a new interview record to the database
	// @Tags Interviews
	// @Accept json
	// @Produce json
	// @Param request body domain.Interview true "Interview Data"
	// @Success 201 {object} map[string]string "Interview created successfully"
	// @Failure 400 {object} map[string]string "Invalid request"
	// @Failure 500 {object} map[string]string "Internal server error"
	// @Router /interviews [post]
	r.POST("/interviews", jwtUtil.AuthMiddleware(cfg.JWTSecretKey), handler.CreateInterview)

	// Health check route
	// @Summary Health Check
	// @Description Returns the health status of the service
	// @Tags Health
	// @Produce json
	// @Success 200 {object} map[string]string "Service is healthy"
	// @Router /health [get]
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "healthy"}) })

	log.Printf("Interview Service is running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
