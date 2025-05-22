package model

// 存放从http请求中取出的参数

// PageParams: 首页组件发送的分页参数
// Page： 页码
// Size： 每页数量
type PageParams struct {
	Page int `json:"page" form:"page"`
	Size int `json:"size" form:"size"`
}
