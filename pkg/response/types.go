package response

type PaginationMeta struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	TotalPage   int `json:"total_page"`
	TotalData   int `json:"total_data"`
}

type ListData struct {
	List       interface{}    `json:"list"`
	Pagination PaginationMeta `json:"pagination"`
}

// @Description Success response with data
type SuccessResponse struct {
	Success bool        `json:"success" example:"true"`
	Data    interface{} `json:"data"`
}

// @Description Success response with message
type MessageResponse struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"device deleted"`
}

// @Description Error response
type ErrorResponse struct {
	Success bool              `json:"success" example:"false"`
	Code    string            `json:"code" example:"ERROR_CODE"`
	Message string            `json:"message" example:"error message"`
	Details map[string]string `json:"details,omitempty" swaggerignore:"true"`
}
