package response

import "time"

type Base struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Timestamp  time.Time   `json:"timestamp"`
	Data       interface{} `json:"data"`
}

type Pagination struct {
	Total       int `json:"total"`
	PerPage     int `json:"perPage"`
	CurrentPage int `json:"currentPage"`
	LastPage    int `json:"lastPage"`
}
