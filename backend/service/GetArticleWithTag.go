package service

import (
	"github.com/zhenqiiii/BlogSite/backend/gorm/sqlite"
	"github.com/zhenqiiii/BlogSite/backend/model"
)

// GetArticleWithTag: 类别文章获取函数,同样需要分页
func GetArticleWithTag(tag string, p model.PageParams) (list []model.ArticleParams, count int, err error) {
	// 获取文章数组：分页查询
	list, err = sqlite.GetArticleByTag(tag, p)
	if err != nil {
		return nil, 0, err
	}
	// 获取分类下的文章总数
	count, err = sqlite.GetTagCount(tag)
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil

}
