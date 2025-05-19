package pkg

import (
	"fiber-template/config"
	"github.com/sirupsen/logrus"
)

const (
	CodeOk       = 0  // 成功
	CodeErr      = -1 // 失败
	CodeErrToken = -2 // token相关的异常
)

// DataResponse represents an HTTP response which contains a JSON body.
type DataResponse struct {
	Code int         `json:"code"` // 0: 成功, -1: 失败, -2: token相关的异常
	Data interface{} `json:"data"` // json数据
}

func SuccessResponse(data interface{}) DataResponse {
	return DataResponse{
		Code: 0,
		Data: data,
	}
}

// MessageResponse returns a JSONResponse with a 'message' key containing the given text.
func MessageResponse(code int, msg, msgZh string) DataResponse {
	config.Log.WithFields(logrus.Fields{
		"code":   code,
		"msg_zh": msgZh,
	}).Warnf(msg)
	return DataResponse{
		Code: code,
		Data: struct {
			Message   string `json:"message"`
			MessageZh string `json:"message_zh"`
		}{msg, msgZh},
	}
}
