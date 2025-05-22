package sqlite

import "github.com/zhenqiiii/BlogSite/backend/model"

// Get: 查询函数集合

// GetList:分页查询文章
func GetList(page int, size int) ([]model.ArticleParams, error) {
	// Limit & Offset ,此外GORM会自动匹配字段
	// Offset为0
	list := make([]model.ArticleParams, 0)
	result := db.Model(&model.Article{}).Limit(size).Offset((page) * size).Find(&list)
	if err := result.Error; err != nil {
		return nil, err
	}
	return list, nil
}

// GetTotal:获取文章总数
func GetTotal() (int, error) {
	// 查询总数
	result := db.Find(&model.Article{})
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil

}
