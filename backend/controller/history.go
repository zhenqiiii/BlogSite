package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/BlogSite/backend/constants"
	"github.com/zhenqiiii/BlogSite/backend/service"
	"github.com/zhenqiiii/BlogSite/backend/utils"
)

// HistoryHandler: 历史路由处理函数
// 功能：接收History组件的请求，返回网站的上传历史数据
// 说明：直接从article表中取标题和date数据，排序由前端还是后端完成？
func HistoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 无参数
		// 业务逻辑
		list, err := service.GetHistory()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": constants.InternalError,
				"msg":  "Something went wrong with GetHistory",
				"err":  err.Error(),
			})
		}
		// 返回数据
		utils.SuccessResponse(c, "Get History Successfully", list)
	}
}
