package model

import (
	"Gin-Blog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func (Category) TableName() string {
	return "category"
}

/*
查询分类是否存在,返回code
*/
func CheckCategory(name string) (code int) {
	var cate Category
	// 使用了 GORM 来查询数据库
	db.Select("id").Where("name = ?", name).First(&cate)
	// 用户是否存在
	if cate.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED
	}
	return errmsg.SUCCESS
}

/*
新增分类, 返回code，判断是否成功
*/
func CreateCate(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

/*
查询分类下的所有文章
*/

/*
查询分类列表，分页查询
pageSize：每页显示的记录数
pageNum：当前页码
返回 []User 类型的切片
*/
func GetCate(pageSize int, pageNum int) ([]Category, int64) {
	var cate []Category
	var total int64
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

/*
编辑分类：只能编辑username和role
*/
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

/*
删除分类：软删除
*/
func DeleteCate(id int) int {
	var cate Category
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
