package utils

import (
	"boilerplate/constant"
	"math"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetPaginationParams(c echo.Context) (int, int) {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = constant.DefaultPage
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 || limit > constant.MaxLimit {
		limit = constant.DefaultLimit
	}

	return page, limit
}

func CreateMeta(page, limit int, total int64) *Meta {
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &Meta{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}
}
