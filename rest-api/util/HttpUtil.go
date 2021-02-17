package util

import "net/http"

type BaseResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BaseResponse(code int, data interface{}) BaseResponseData {
	return BaseResponseData{
		Code:    code,
		Message: http.StatusText(code),
		Data:    data,
	}
}
