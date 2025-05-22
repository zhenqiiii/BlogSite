package sqlite

import (
	"fmt"

	"github.com/zhenqiiii/BlogSite/backend/model"
)

// 存放创建相关函数

// CreateArticle:插入文章函数
func CreateArticle(article model.Article) error {
	result := db.Create(&article)
	if result.Error != nil {
		return result.Error
	}
	fmt.Print("插入成功")
	return nil
}
