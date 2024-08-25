package model

import (
	"Gin-Blog/utils/errmsg"
	"gorm.io/gorm"
)

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

/*
新增文章, 返回code，判断是否成功
*/
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

/*
查询分类下的所有文章
因为有多条文章，所以需要分页处理
id : 分类id
pageSize：每页显示的记录数
pageNum：当前页码
返回 []Article 类型的切片
*/
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&cateArtList).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCESS, total
}

/*
查询单个文章
*/
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

/*
查询文章列表，分页查询
pageSize：每页显示的记录数
pageNum：当前页码
返回 []Article 类型的切片
*/
func GetArt(pageSize int, pageNum int) ([]Article, int, int64) {
	var aticleList []Article
	var total int64
	// 查找文章时 预加载Category
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&aticleList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return aticleList, errmsg.SUCCESS, total
}

/*
编辑文章：只能编辑username和role
*/
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err := db.Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

/*
删除文章：软删除
*/
func DeleteArt(id int) int {
	var art Article
	err := db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
