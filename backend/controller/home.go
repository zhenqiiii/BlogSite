package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/BlogSite/backend/constants"
	"github.com/zhenqiiii/BlogSite/backend/model"
	"github.com/zhenqiiii/BlogSite/backend/service"
	"github.com/zhenqiiii/BlogSite/backend/utils"
)

// HomeHandler: 主页路由处理函数
// 功能：接收到Home组件的请求后返回一个文章数组&文章总数
// 说明：前端发送请求时附带页码，后端执行分页查询
func HomeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 处理页码参数
		var p model.PageParams
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": constants.FaultParams,
				"msg":  "Check the params again",
				"err":  err,
			})
			return
		}
		// 执行业务逻辑: 分页查询&总数查询
		list, total, err := service.HomeArticleList(p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": constants.InternalError,
				"msg":  "Internal Server Error",
				"err":  err,
			})
			return
		}

		// 返回数据：article-list & total
		utils.SuccessResponse(c, "Success", gin.H{"total": total, "list": list})

	}

}
