package model

import "gorm.io/gorm"

// 实现 Tabler 接口来更改默认复数表名为单数
type Tabler interface {
	TableName() string
}

/*
定义了一个名为 User 的 Go 语言结构体，
用于表示一个数据库中的用户表。结构体中的
字段被绑定到数据库的列，并且使用了 GORM
提供的标签来指定字段的属性。
*/
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// 将默认的 users 表名重写为单数 user
func (User) TableName() string {
	return "user"
}
