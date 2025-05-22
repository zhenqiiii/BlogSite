package service

import (
	"github.com/zhenqiiii/BlogSite/backend/gorm/sqlite"
	"github.com/zhenqiiii/BlogSite/backend/model"
)

// HomeArticleList: Home组件获取文章列表业务函数
// 接收PageParams参数，返回err & 文章数组
// 暂时不确定后续能不能复用
func HomeArticleList(p model.PageParams) (article_list []model.ArticleParams, total int, err error) {
	// 执行分页查询
	article_list, err = sqlite.GetList(p.Page, p.Size)
	if err != nil {
		return nil, 0, err
	}
	// 获取total
	total, err = sqlite.GetTotal()
	if err != nil {
		return nil, 0, err
	}

	return article_list, total, nil

}
