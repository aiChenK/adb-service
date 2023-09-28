package routes

import (
	"adb-service/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoute(engine *gin.Engine) {

	engine.Any(
		"/adb",
		func(c *gin.Context) {
			(&controllers.AdbController{}).Index(c)
		},
	)
}
