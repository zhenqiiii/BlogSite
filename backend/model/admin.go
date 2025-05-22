package model

type Admin struct {
	ID        int64  `gorm:"primaryKey"`
	AdminName string `json:"adminname" gorm:"type:varchar(100)"`
	Password  string `json:"password" form:"password" gorm:"type:varchar(200)"`
}
