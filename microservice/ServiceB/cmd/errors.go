package cmd

import (
	"fmt"
	"runtime"
	"time"
)

// 定义一个结构体来表示错误信息

type CustomError struct {
	Message   string    // 错误信息
	File      string    // 错误发生的文件名
	Line      int       // 错误发生的行号
	Timestamp time.Time // 错误发生的时间
}

// 捕获错误的函数

func NewError(message string) *CustomError {
	// 获取调用者的信息
	_, file, line, ok := runtime.Caller(1) // 1 表示获取上一级函数调用的堆栈信息

	if !ok {
		file = "unknown"
		line = 0
	}

	return &CustomError{
		Message:   message,
		File:      file,
		Line:      line,
		Timestamp: time.Now(),
	}
}

// 打印错误信息
func (e *CustomError) Error() string {
	return fmt.Sprintf("[%s] Error: %s (File: %s, Line: %d)", e.Timestamp.Format(time.RFC3339), e.Message, e.File, e.Line)
}
