package controllers

import (
	"adb-service/core"
	"adb-service/dto/request"
	"github.com/gin-gonic/gin"
	"net/http"
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
	if (commandForm.Cmd == "setProxy") && commandForm.ProxyAddr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":    false,
			"errMessage": "包名不可为空",
		})
		return
	}
	if (commandForm.Cmd == "clear" || commandForm.Cmd == "stop") && commandForm.PackageName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":    false,
			"errMessage": "包名不可为空",
		})
		return
	}
	if commandForm.Port == "" {
		commandForm.Port = "5555"
	}

	err = core.Exec(commandForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":    false,
			"errMessage": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":    true,
		"errMessage": "执行成功",
	})
}
