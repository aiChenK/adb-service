package controllers

import (
	"adb-service/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ToolController struct{}

func (c *ToolController) Ip(ctx *gin.Context) {
	// 获取ip
	ip, err := core.GetLocalIp()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success":    false,
			"errMessage": err.Error(),
		})
	}
	// 返回ip
	ctx.JSON(http.StatusOK, gin.H{
		"success":    true,
		"errMessage": "",
		"data":       ip,
	})
}
