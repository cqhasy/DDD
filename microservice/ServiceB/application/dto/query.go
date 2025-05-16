package dto

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
)

//解析请求的相关参数。

func Bind(c *gin.Context, req interface{}) error {
	var err error
	// 根据请求方法选择合适的绑定方式
	if c.Request.Method == http.MethodGet {
		err = c.ShouldBindQuery(req) // 处理GET请求的查询参数
	} else {
		err = c.ShouldBind(req) // 处理POST、PUT等请求的请求体数据
	}
	if err != nil {
		c.Error(err)
		return err
	}
	for _, param := range c.Params {
		// 使用反射动态将路径参数赋值到 req 对象中
		err = bindPathParamToStructField(req, param.Key, param.Value)
		if err != nil {
			return err
		}
	}

	return nil
}
func bindPathParamToStructField(req interface{}, paramName string, paramValue string) error {
	// 获取 req 的反射对象
	v := reflect.ValueOf(req).Elem() // 获取结构体的值（非指针）

	// 确保 req 是指针类型
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("req must be a pointer to a struct")
	}

	// 查找结构体中是否有与路径参数名匹配的字段
	field := v.FieldByName(paramName)
	if !field.IsValid() {
		return nil // 如果没有找到匹配的字段，不进行处理
	}

	// 检查字段是否可以写入
	if !field.CanSet() {
		return fmt.Errorf("cannot set field %s", paramName)
	}

	// 确定字段类型并将值转换为适当类型
	switch field.Kind() {
	case reflect.String:
		field.SetString(paramValue)
	case reflect.Int:
		// 转换为整数并设置
		if intValue, err := strconv.Atoi(paramValue); err == nil {
			field.SetInt(int64(intValue))
		} else {
			return fmt.Errorf("invalid integer value for field %s", paramName)
		}
	case reflect.Bool:
		// 转换为布尔值并设置
		if boolValue, err := strconv.ParseBool(paramValue); err == nil {
			field.SetBool(boolValue)
		} else {
			return fmt.Errorf("invalid boolean value for field %s", paramName)
		}
	default:
		return fmt.Errorf("unsupported field type: %s", field.Kind())
	}

	return nil
}
