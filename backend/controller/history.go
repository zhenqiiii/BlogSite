package controller

import "github.com/gin-gonic/gin"

// HistoryHandler: 历史路由处理函数
// 功能：接收History组件的请求，返回网站的上传历史数据
// 说明：直接从article表中取标题和date数据，排序由前端还是后端完成？
func HistoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 无参数
		// 业务逻辑
		// 返回数据
	}
}
