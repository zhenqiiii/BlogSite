package service

import (
	"github.com/zhenqiiii/BlogSite/backend/gorm/sqlite"
	"github.com/zhenqiiii/BlogSite/backend/model"
)

// GetArticle: 获取文章内容业务函数
func GetArticle(id int64) (article *model.Article, err error) {
	// 查询数据库
	article, err = sqlite.GetArticleByID(id)
	if err != nil {
		return nil, err
	}
	// 返回数据
	return article, nil
}
