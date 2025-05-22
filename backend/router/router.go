package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/BlogSite/backend/controller"
)

// 创建路由实例
func SetupRouter(mode string) *gin.Engine {
	// 模式设置,默认情况下DebugMode
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建实例
	r := gin.Default()

	// 按照前端的路由设置来
	// home路由
	r.GET("/home", controller.HomeHandler())

	//history路由
	r.GET("/history", controller.HistoryHandler())

	// article路由
	r.GET("/article/:id", controller.ArticleHandler())

	// tag 路由
	r.GET("/tag/:tag", controller.TagHandler())

	// admin:后台管理路由,需要鉴权
	adminGroup := r.Group("/admin")
	{
		// upload:上传路由,POST,在upload页面点击上传按钮后发送请求
		adminGroup.POST("/upload", controller.UploadHandler())

		// management: 文章管理路由-查看/编辑/删除
		managGroup := r.Group("/management")
		{
			// 查看某篇文章
			managGroup.GET("/article/:id", controller.ArticleHandler())
			// 编辑文章内容: 点击查看页面的编辑按钮后进入修改界面，然后点击提交按钮发送请求
			managGroup.POST("/edit", controller.EditHandler())
			// 删除文章：同样点击查看后在查看页面进行删除操作
			managGroup.GET("/delete/:id", controller.DeleteHandler())

		}
	}

	// 无此路径时
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "No such route",
		})
	})

	return r
}
