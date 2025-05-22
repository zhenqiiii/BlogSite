package controller

import "github.com/gin-gonic/gin"

// ArticleHandler: 文章查看处理函数
// 功能：接收Article组件的请求，根据附带的id查询并返回文章内容
// 说明：在首页点击文章标题后进入Article组件，组件再向后端发送请求
// Article组件似乎可以在管理界面中的文章查看页面复用，因此该处理函数也可以复用
func ArticleHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接收文章id参数
		// 业务逻辑:查询并获取
		// 返回数据
	}

}
