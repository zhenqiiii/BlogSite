package service

import (
	"github.com/zhenqiiii/BlogSite/backend/gorm/sqlite"
	"github.com/zhenqiiii/BlogSite/backend/model"
)

// GetHistory:获取历史上传数组的业务函数
func GetHistory() (list []model.DateParams, err error) {
	// 查询数据库
	list, err = sqlite.GetHistory()
	if err != nil {
		return nil, err
	}
	// 根据CreatedAt字段排序:最早-最晚，日期早的排在前面
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			// 若后面的比前面的早，交换位置
			if list[j].CreatedAt.Before(list[i].CreatedAt) {
				tmp := list[i]
				list[i] = list[j]
				list[j] = tmp
			}
		}
	}
	return list, nil
}
