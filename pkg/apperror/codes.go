package apperror

import "net/http"

const (
	NotFound         = "NOT_FOUND"
	AlreadyExists    = "ALREADY_EXISTS"
	ValidationFailed = "VALIDATION_FAILED"
	InvalidRequest   = "INVALID_REQUEST"
	InternalError    = "INTERNAL_ERROR"
)

var codeMap = map[string]int{
	NotFound:         http.StatusNotFound,
	AlreadyExists:    http.StatusConflict,
	ValidationFailed: http.StatusBadRequest,
	InvalidRequest:   http.StatusBadRequest,
	InternalError:    http.StatusInternalServerError,
}

func HTTPStatus(bizCode string) int {
	if code, ok := codeMap[bizCode]; ok {
		return code
	}
	return http.StatusInternalServerError
}
