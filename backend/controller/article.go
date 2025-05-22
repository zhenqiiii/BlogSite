package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/BlogSite/backend/constants"
	"github.com/zhenqiiii/BlogSite/backend/service"
	"github.com/zhenqiiii/BlogSite/backend/utils"
)

// ArticleHandler: 文章查看处理函数
// 功能：接收Article组件的请求，根据附带的id查询并返回文章内容
// 说明：在首页点击文章标题后进入Article组件，组件再向后端发送请求
// Article组件似乎可以在管理界面中的文章查看页面复用，因此该处理函数也可以复用
func ArticleHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接收文章id参数:获取路径参数并转int64
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		// 业务逻辑:查询并获取
		article, err := service.GetArticleWithID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": constants.InternalError,
				"msg":  "Fail to get article",
				"err":  err.Error(),
			})
			return
		}
		// 返回数据
		utils.SuccessResponse(c, "Get article successfully", article)
	}

}
