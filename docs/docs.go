// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://example.com/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.example.com/support",
            "email": "support@example.com"
        },
        "license": {
            "name": "MIT License",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "Returns the health status of the interview service",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Check service health",
                "responses": {
                    "200": {
                        "description": "Service is healthy",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/interviews": {
            "get": {
                "description": "Retrieve a list of all interviews in the system",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Interviews"
                ],
                "summary": "Get all interviews",
                "responses": {
                    "200": {
                        "description": "List of interviews",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Interview"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to fetch interviews",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new interview by providing candidate_id, job_id, interview_date, and feedback",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Interviews"
                ],
                "summary": "Create a new interview",
                "parameters": [
                    {
                        "description": "Interview Creation Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Interview"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Interview created successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to create interview",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Interview": {
            "type": "object",
            "properties": {
                "candidate_id": {
                    "description": "Foreign key referencing the candidate's ID",
                    "type": "integer"
                },
                "feedback": {
                    "description": "Feedback or notes about the interview",
                    "type": "string"
                },
                "id": {
                    "description": "Unique identifier for the interview",
                    "type": "integer"
                },
                "interview_date": {
                    "description": "Date and time of the interview",
                    "type": "string"
                },
                "job_id": {
                    "description": "Foreign key referencing the job's ID",
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Interview Service API",
	Description:      "API for managing interviews in the system.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
