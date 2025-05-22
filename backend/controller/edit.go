package controller

import "github.com/gin-gonic/gin"

// EditHandler: 编辑文章路由处理函数
// 功能：接收修改后的内容,对已经存在的内容进行覆盖
// 说明: 在查看界面点击编辑按钮后，跳转至编辑器界面，完成内容修改后，点击完成按钮后向该API发送请求
func EditHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接收数据并创建实例
		// 业务逻辑:原有内容进行覆盖
		// 返回响应
	}
}
