package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/BlogSite/backend/constants"
	"github.com/zhenqiiii/BlogSite/backend/model"
	"github.com/zhenqiiii/BlogSite/backend/service"
	"github.com/zhenqiiii/BlogSite/backend/utils"
)

// UploadHandler： 上传路由处理函数
// 功能： 接收要上传的文章数据，存入数据库表
// 说明： 在Admin组件中存在一个上传页面入口，进入后在上传页面编辑内容，完成后点击上传按钮（此处发送请求）
func UploadHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接收参数，创建实例
		// 参数:标题/摘要/内容/标签
		var article model.Article
		err := c.ShouldBind(&article)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": constants.FaultParams,
				"msg":  "check the params again",
				"err":  err,
			})
			return
		}
		// 业务逻辑：存入
		err = service.UploadArticle(article)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": constants.InternalError,
				"msg":  "UploadArticle method",
				"err":  err,
			})
			return
		}
		// 返回响应
		utils.SuccessResponse(c, "Upload Successfully", nil)
	}
}
