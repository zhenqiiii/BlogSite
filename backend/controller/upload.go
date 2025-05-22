package controller

import "github.com/gin-gonic/gin"

// UploadHandler： 上传路由处理函数
// 功能： 接收要上传的文章数据，存入数据库表
// 说明： 在Admin组件中存在一个上传页面入口，进入后在上传页面编辑内容，完成后点击上传按钮（此处发送请求）
func UploadHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接收参数，创建实例
		// 业务逻辑：存入
		// 返回响应
	}
}
