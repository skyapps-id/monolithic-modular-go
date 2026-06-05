package response

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func ParsePagination(c echo.Context) (page, limit int) {
	page, _ = strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	limit, _ = strconv.Atoi(c.QueryParam("limit"))
	if limit < 1 {
		limit = 12
	}
	return
}

func NewPaginationMeta(page, limit, total int) PaginationMeta {
	totalPage := 0
	if total > 0 {
		totalPage = (total + limit - 1) / limit
	}
	return PaginationMeta{
		CurrentPage: page,
		PerPage:     limit,
		TotalPage:   totalPage,
		TotalData:   total,
	}
}
