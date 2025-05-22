package utils

import "math/rand/v2"

// GenArticleID: 文章ID生成函数
func GenArticleID() (ArticleID int64) {
	ArticleID = rand.Int64N(100000) + 100000
	return ArticleID
}
