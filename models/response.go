package models

import "encoding/json"

type Err interface {
	Code() int
	Message() string
	IsSuccessful() bool
}

type errResp struct {
	ECode    int    `json:"code"`
	EMessage string `json:"msg"`
}

func (e *errResp) Code() int {
	return e.ECode
}

func (e *errResp) Message() string {
	return e.EMessage
}

func (e *errResp) IsSuccessful() bool {
	return e.ECode == SuccessCode || e.ECode == CodeMsgExisted
}

func (e *errResp) Error() string {
	bs, err := json.Marshal(e)
	if err != nil {
		return "{}"
	}

	return string(bs)
}

func NewErr(code int, message string) Err {
	return &errResp{
		ECode:    code,
		EMessage: message,
	}
}

func NewErrWithBytes(bytes []byte) Err {
	var e errResp
	err := json.Unmarshal(bytes, &e)
	if err != nil {
		return &errResp{
			ECode:    -1,
			EMessage: string(bytes),
		}
	}

	return &e
}

type OkResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

const (
	SuccessCode    = 0
	CodeMsgExisted = 101002
)
