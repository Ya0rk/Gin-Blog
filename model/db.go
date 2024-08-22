package model

import (
	"Gin-Blog/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 数据库链接的参数

var db *gorm.DB
var err error

func InitDb() {
	// 打开数据库
	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.Dbname,
	)), &gorm.Config{})

	if err != nil {
		fmt.Println("连接数据库失败：", err)
	}

	// 数据库迁移, 创建这三张表
	db.AutoMigrate(&User{}, &Article{}, &Category{})

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("获取通用数据库对象 sql.DB 出错：", err)
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	// 关闭数据库
	//sqlDB.Close()
}
