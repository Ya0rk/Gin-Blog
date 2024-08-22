package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`
	Desc     string   `gorm:"type:text;" json:"desc"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Img      string   `gorm:"type:varchar(100)" json:"img"`
	Category Category `gorm:"foreignkey:Cid;references:ID"` // 修正的外键定义
}

func (Article) TableName() string {
	return "article"
}
