package sqlite

import (
	"github.com/zhenqiiii/BlogSite/backend/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 全局定义连接实例
var db *gorm.DB

// ConnDB:数据库连接函数
func ConnDB() (err error) {
	db, err = gorm.Open(sqlite.Open("Blog.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	// 自动创建/更新表
	db.AutoMigrate(&model.Article{}, &model.Admin{})

	return nil
}
