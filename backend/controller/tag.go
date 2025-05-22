package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/BlogSite/backend/constants"
	"github.com/zhenqiiii/BlogSite/backend/model"
	"github.com/zhenqiiii/BlogSite/backend/service"
	"github.com/zhenqiiii/BlogSite/backend/utils"
)

// TagHandler: 标签(分类)处理函数
// 功能: 接收请求,根据请求附带的tag和page参数,返回对应tag的文章
// 说明： 类似于首页的文章展示功能，也需要实现分页
func TagHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接收Tag和页码参数
		var p model.PageParams
		tag := c.Param("tag")
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": constants.FaultParams,
				"msg":  "Check the params again",
				"err":  err.Error(),
			})
			return
		}
		// 业务逻辑：查询对应文章,返回文章数组
		list, count, err := service.GetArticleWithTag(tag, p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": constants.InternalError,
				"msg":  "something went wrong with GetArticleWithTag",
				"err":  err.Error(),
			})
			return
		}
		// 返回文章数组
		utils.SuccessResponse(c, "Get "+tag+" article successfully", gin.H{
			"list":  list,
			"total": count,
		})
	}

}
