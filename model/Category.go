package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func (Category) TableName() string {
	return "category"
}
