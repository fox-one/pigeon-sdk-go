package models

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func (e *ErrorResponse) IsSuccessful() bool {
	return e.Code == SuccessCode || e.Code == CodeMsgExisted
}

type OkResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

const (
	SuccessCode    = 0
	CodeMsgExisted = 101002
)
