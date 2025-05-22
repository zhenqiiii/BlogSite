package main

import (
	"fmt"

	"github.com/zhenqiiii/BlogSite/backend/gorm/sqlite"
	"github.com/zhenqiiii/BlogSite/backend/router"
)

func main() {
	// 连接数据库
	err := sqlite.ConnDB()
	if err != nil {
		fmt.Printf("连接数据库失败:%v", err)
	}
	// 初始化路由
	r := router.SetupRouter("")

	// 这次是大🐎
	err = r.Run(":8080")
	if err != nil {
		fmt.Printf("服务器启动失败：%v", err)
	}
}
