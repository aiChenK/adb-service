package routes

import (
	"adb-service/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(engine *gin.Engine) {

	// 添加跨域中间件
	engine.Use(CORSMiddleware())

	engine.Any(
		"/adb",
		func(c *gin.Context) {
			(&controllers.AdbController{}).Index(c)
		},
	)
}

// CORSMiddleware 是跨域中间件函数
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许跨域的域名，如果 * 则允许所有域名
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 设置允许的请求方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		// 设置允许的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// 如果请求方法是OPTIONS，则直接返回，不再继续处理
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		// 继续处理请求
		c.Next()
	}
}
