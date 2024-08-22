package main

import (
	"Gin-Blog/model"
	"Gin-Blog/routes"
)

func main() {
	// 引用数据库
	model.InitDb()
	routes.InitRouter()
}
