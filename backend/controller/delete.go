package controller

import "github.com/gin-gonic/gin"

// DeleteHandler:删除文章路由函数
// 功能: 接收删除请求,执行删除操作
// 说明: 请求由查看页面的删除按钮触发,请求携带对应文章id
func DeleteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接收id参数
		// 业务逻辑:查询并删除
		// 返回响应

	}
}
