package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/BlogSite/backend/constants"
)

// 成功响应封装函数
func SuccessResponse(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": constants.Success,
		"msg":  msg,
		"data": data,
	})
}

// // 错误响应函数
// func FailResponse (c *gin.Context, msg string, err error) {
// 	c.JSON()
// }
