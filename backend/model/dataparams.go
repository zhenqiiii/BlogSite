package model

import "time"

// 存放data数据参数,即响应参数

// ArticleParams : 向首页组件发送的文章
// 说明:由于首页不需要某些字段,所以从Article模型中剥离这些字段
type ArticleParams struct {
	Title     string    `json:"title" form:"title"`
	Digest    string    `json:"digest" form:"digest"`
	Tag       string    `json:"tag" form:"tag"`
	ReadCount int       `json:"readcount" form:"readcount"`
	CreatedAt time.Time `json:"createdat" form:"createdat"`
}
