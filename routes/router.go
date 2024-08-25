package routes

import (
	"Gin-Blog/api/v1"
	"Gin-Blog/middleware"
	"Gin-Blog/utils"
	"github.com/gin-gonic/gin"
)

/*
	创建路由
*/

func InitRouter() {
	// 先gin.SetMode()，然后gin.New或者gin.Default()，配置才能生效
	gin.SetMode(utils.AppMode)
	r := gin.New() // 不用自带的中间件， 选择自行添加
	r.Use(middleware.Logger())
	r.Use(gin.Recovery()) // 手动添加 Recovery 中间件

	// 创建路由组管理
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// User模块的路由接口
		auth.PUT("users/:id", v1.EditUser)
		auth.DELETE("users/:id", v1.DeleteUser)
		// Category模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)

		// Article模块的路由接口
		auth.POST("article/add", v1.AddArt)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)

		// 上传文件
		auth.POST("upload", v1.Upload)
	}

	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser) // 新增用户

		router.GET("users", v1.GetUsers)
		router.GET("category", v1.GetCate)

		router.GET("article", v1.GetArt)              // 查询文章列表
		router.GET("article/list/:id", v1.GetCateArt) // 查询分类下的所有文章
		router.GET("article/info/:id", v1.GetArtInfo) // 查询单个文章信息
		router.POST("login", v1.Login)
	}
	r.Run(utils.HttpPort)
}
