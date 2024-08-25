package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

/*
跨域请求，参数配置
*/
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowMethods:    []string{"*"},
			AllowHeaders:    []string{"Origin"},
			ExposeHeaders:   []string{"Content-Length", "Authorization"},
			//AllowCredentials: true,
			MaxAge: 12 * time.Hour,
		})
	}
}
