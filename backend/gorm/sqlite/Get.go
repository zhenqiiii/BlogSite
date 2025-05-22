package sqlite

import (
	"fmt"

	"github.com/zhenqiiii/BlogSite/backend/model"
)

// Get: 查询函数集合

// GetList:分页查询文章
func GetList(page int, size int) (list []model.ArticleParams, err error) {
	// Limit & Offset ,此外GORM会自动匹配字段
	// Offset为0
	result := db.Model(&model.Article{}).Limit(size).Offset((page - 1) * size).Find(&list)
	fmt.Print(result.RowsAffected)
	if err = result.Error; err != nil {
		return nil, err
	}
	fmt.Print(list)
	return list, nil
}

// GetTotal:获取文章总数
func GetTotal() (int, error) {
	// 查询总数
	var articles []model.Article
	result := db.Find(&articles)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil

}
