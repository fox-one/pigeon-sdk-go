package models

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type OkResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

const (
	SuccessCode = 0
)

const (
	
)
