package dto

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//todo: 处理日志的方式不太好，小藕合，需改进。

type Response struct {
	Status  string      `json:"status"`  // success 或 error
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 响应信息
	Data    interface{} `json:"data"`    // 数据，成功时返回数据，失败时返回错误信息
}

//通用的返回响应函数。

func Respond(c *gin.Context, data interface{}, err error, message string) {

	if err != nil {
		code := GetErrorCode(err)
		c.JSON(code, Response{
			Status:  "error",
			Code:    code,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})

}
