package routers

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"go.translation.api/extension/dbLog"
	"go.translation.api/extension/server"
	"go.translation.api/http/admin/controller/log.controller"
	"go.translation.api/http/controller/translationController"
	"os/exec"
)


func Routers(){
	//设置静态文件路口
	server.Server.Static("/assets", "./assets")
	version :=server.Server.Group("version")
	{
		version.GET("/", func(context *gin.Context) {
			cmd := exec.Command("/bin/bash", "-c", `git log`)
			//创建获取命令输出管道
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				dbLog.ServerError("Error:can not obtain stdout pipe for command",err)
				context.JSON(-1,gin.H{
					"meeeage":"fail",
				})
				return
			}

			//执行命令
			if err := cmd.Start(); err != nil {
				dbLog.ServerError("Error:The command is err",err)
				context.JSON(-1,gin.H{
					"meeeage":"fail",
				})
			}
			//使用带缓冲的读取器
			outputBuf := bufio.NewReader(stdout)
			//只读区第一行
			output, _, err := outputBuf.ReadLine()
			commit:=string(output)
			context.JSON(200,gin.H{
				"meeeage":"ok",
				"version":commit,
			})
		})
	}

	v1:= server.Server.Group("v1")
	translationGroup:=v1.Group("tran")
	{
		tran:=  translationController.TranslationController{}
		translationGroup.Any("action",tran.ActionTranslation)
		translationGroup.Any("tran_list",tran.ActionTranslationList)
	}



	ser :=server.Server
	ser.LoadHTMLGlob("templates/*")
	web:=ser.Group("web",gin.BasicAuth(gin.Accounts{
		"admin@admin.com":"app@2021",
	}))
	logGroup :=web.Group("log")
	{
		log :=log_controller.LogController{}
		logGroup.GET("list",log.LogList)
	}
}
