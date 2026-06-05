package apperror

import "fmt"

type AppError struct {
	HTTPCode int               `json:"-"`
	BizCode  string            `json:"code"`
	Message  string            `json:"message"`
	Details  map[string]string `json:"details,omitempty"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s", e.BizCode, e.Message)
}

func (e *AppError) WithDetails(details map[string]string) *AppError {
	e.Details = details
	return e
}

func New(bizCode string, message string) *AppError {
	return &AppError{HTTPCode: HTTPStatus(bizCode), BizCode: bizCode, Message: message}
}

func NotFoundError(entity string) *AppError {
	return New(NotFound, entity+" not found")
}

func AlreadyExistsError(entity string) *AppError {
	return New(AlreadyExists, entity+" already exists")
}

func ValidationError(message string) *AppError {
	return New(ValidationFailed, message)
}
