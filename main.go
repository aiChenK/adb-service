package main

import (
	"adb-service/core"
	"adb-service/routes"
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed static/dist
var Static embed.FS

func main() {
	// 检查adb命令
	core.CheckAdb()

	// 创建默认路由引擎
	r := gin.Default()
	routes.InitRoute(r)

	// 静态资源(中间件避免路由冲突)
	r.Use(routes.Serve("/", routes.EmbedFolder(Static, "static/dist")))
	r.NoRoute(func(c *gin.Context) {
		data, err := Static.ReadFile("static/dist/index.html")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})

	port := "8088"
	go func() {
		// 打开浏览器
		core.OpenBrowser("http://localhost:" + port)
	}()

	// 启动
	r.Run(":" + port)
}
