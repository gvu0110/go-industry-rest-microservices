package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

// APIError interface
type APIError interface {
	GetStatusCode() int
	GetMessage() string
	GetError() string
}

type apiError struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Error      string `json:"error,omitempty"`
}

func (e *apiError) GetStatusCode() int {
	return e.StatusCode
}

func (e *apiError) GetMessage() string {
	return e.Message
}

func (e *apiError) GetError() string {
	return e.Error
}

// NewAPIError function
func NewAPIError(statusCode int, message string) APIError {
	return &apiError{
		StatusCode: statusCode,
		Message:    message,
	}
}

// NewNotFoundError function
func NewNotFoundError(message string) APIError {
	return &apiError{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

// NewInternalServerError function
func NewInternalServerError(message string) APIError {
	return &apiError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}

// NewBadRequestError function
func NewBadRequestError(message string) APIError {
	return &apiError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

// NewAPIErrorFromBytes function
func NewAPIErrorFromBytes(body []byte) (APIError, error) {
	var result apiError
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("Invalid JSON body")
	}
	return &result, nil
}
