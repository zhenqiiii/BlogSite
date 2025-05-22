package sqlite

import (
	"errors"
	"fmt"

	"github.com/zhenqiiii/BlogSite/backend/model"
	"gorm.io/gorm"
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

// GetArticleByID：通过id获取整篇文章
func GetArticleByID(id int64) (article *model.Article, err error) {
	// 内联条件
	// 如果没查到该id的文章，由于GORM查询时添加了LIMIT 1条件，所以会报错：result.Error

	result := db.Where("p_id = ?", id).First(&article)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {

	}
	if err = result.Error; err != nil {
		return nil, err
	}
	return article, nil

}
