package dto

import (
	"ServiceB/cmd"
	"log"
	"net/http"
)

//todo:在这里自定义要最终返回的哨兵错误，并于相应的状态码建立连接。

var ErrorCodeMap = map[string]int{
	"DatabaseError":     http.StatusInternalServerError,
	"ValidationError":   http.StatusBadRequest,
	"NotFoundError":     http.StatusNotFound,
	"UnauthorizedError": http.StatusUnauthorized,
}

func GetErrorCode(err error) int {
	// 假设我们根据 err.Error() 来获取错误类型，你可以根据需求自定义
	c := cmd.NewError(err.Error())
	log.Println(c.Error())
	if code, ok := ErrorCodeMap[err.Error()]; ok {
		return code
	}

	// 如果找不到特定类型，默认返回 500
	return http.StatusInternalServerError
}
