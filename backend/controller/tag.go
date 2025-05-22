package controller

import "github.com/gin-gonic/gin"

// TagHandler: 标签(分类)处理函数
// 功能: 接收请求,根据请求附带的tag参数,返回对应tag的文章
// 说明： 类似于首页的文章展示功能，也需要实现分页
func TagHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接收Tag参数
		// 业务逻辑：查询对应文章
		// 返回文章数组
	}

}
