package routes

import (
	"Gin-Blog/api/v1"
	"Gin-Blog/utils"
	"github.com/gin-gonic/gin"
)

/*
	创建路由
*/

func InitRouter() {
	// 先gin.SetMode()，然后gin.New或者gin.Default()，配置才能生效
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	// 创建路由组管理
	router := r.Group("api/v1")
	{
		// User模块的路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("users/:id", v1.EditUser)
		router.DELETE("users/:id", v1.DeleteUser)

		// Category模块的路由接口

		// Article模块的路由接口
	}

	r.Run(utils.HttpPort)
}
