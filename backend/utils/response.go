package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/BlogSite/backend/constants"
)

// 正常响应封装函数
func SuccessResponse(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": constants.Success,
		"msg":  msg,
		"data": data,
	})
}
