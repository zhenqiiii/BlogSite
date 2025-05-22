package model

import "time"

// 存放data数据参数,即响应参数

// ArticleParams : 向首页组件发送的文章
// 说明:由于首页不需要某些字段,所以从Article模型中剥离这些字段
type ArticleParams struct {
	Title     string    `json:"createdat" form:"createdat"`
	Digest    string    `json:"digest" form:"digest"`
	Tag       string    `json:"tag" form:"tag"`
	ReadCount int       `json:"readcount" form:"readcount"`
	CreatedAt time.Time `json:"createdat" form:"createdat"`
}

// DateParams：存放用于渲染History组件的文章数据
type DateParams struct {
	PID       int64     `json:"pid" form:"pid"`
	Title     string    `json:"title" form:"title"`
	CreatedAt time.Time `json:"createdat" form:"createdat"`
}
