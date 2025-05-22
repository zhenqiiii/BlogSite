package service

import (
	"github.com/zhenqiiii/BlogSite/backend/gorm/sqlite"
	"github.com/zhenqiiii/BlogSite/backend/model"
	"github.com/zhenqiiii/BlogSite/backend/utils"
)

// UploadArticle: 上传文章业务函数
func UploadArticle(article model.Article) (err error) {
	// 生成ID
	article.PID = utils.GenArticleID()
	// 插入数据库
	err = sqlite.CreateArticle(article)
	if err != nil {
		return err
	}
	return nil
}
