package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const SuccessCode = http.StatusOK

type ResponseMessage struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (resp *ResponseMessage) ToJson(c *gin.Context) {
	c.JSON(SuccessCode, resp)
}

//错误返回
func ErrResponse(code int, msg string, data interface{}) *ResponseMessage {
	return &ResponseMessage{
		Status:  code,
		Message: msg,
		Data:    data,
	}
}

//成功返回
func SuccessResponse(data interface{}) *ResponseMessage {
	return &ResponseMessage{
		Status:  SuccessCode,
		Message: "Success",
		Data:    data,
	}
}
