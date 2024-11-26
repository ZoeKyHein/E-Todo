package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 数据
}

// Success 成功返回
func Success(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  msg,
		Data: data,
	})
}

// Fail 失败返回
func Fail(c *gin.Context, data interface{}, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
