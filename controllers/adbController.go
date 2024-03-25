package controllers

import (
	"adb-service/core"
	"adb-service/dto/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdbController struct {
}

func (c *AdbController) Index(ctx *gin.Context) {
	commandForm := request.CommandForm{}
	err := ctx.ShouldBind(&commandForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":    false,
			"errMessage": "缺少必要参数",
			"error":      err.Error(),
		})
		return
	}

	// 条件判断
	switch commandForm.Op {
	case "setProxy":
		if commandForm.ProxyAddr == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success":    false,
				"errMessage": "代理地址不可为空",
			})
			return
		}
	case "clear", "stop":
		if commandForm.PackageName == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success":    false,
				"errMessage": "包名不可为空",
			})
			return
		}
	default:
		if commandForm.Op != "free" && commandForm.Ip == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success":    false,
				"errMessage": "客户端ip不可为空",
			})
			return
		}
	}

	// 默认端口
	if commandForm.Port == "" {
		commandForm.Port = "5555"
	}

	// 执行命令
	output, err := core.Exec(commandForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":    false,
			"errMessage": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":    true,
		"errMessage": "",
		"data":       output,
	})
}
