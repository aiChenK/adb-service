package main

import (
	"adb-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建默认路由引擎
	r := gin.Default()

	routes.InitRoute(r)
	// 启动
	r.Run(":8089")

}
