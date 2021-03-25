package routers

import (
	"gin.test/extension/log"
	"gin.test/extension/server"
	"gin.test/http/controller/userController"
	"gin.test/http/middleware"
	"github.com/gin-gonic/gin"
)


func Routers(){
	//设置静态文件路口
	server.Server.Static("/assets", "./assets")


	server.Server.GET("/logs",log.ReadLogs)
	server.Server.GET("/logs_info", log.ReadLogInfo)

	v1:= server.Server.Group("v1")
	//路由组中添加中间件
	//v1.Use(middleware.HttpHeaderVerification())
	v1.Use(middleware.AppMiddleWare())
	{
		v1.GET("/ping", func(context *gin.Context) {
			context.JSON(200,gin.H{
				"meeeage":"ok",
				"id":context.Query("id"),
			})
		})
	}

	//区分路由组
	userGroup :=v1.Group("/user")
	{
		//绑定controller
		user :=userController.UserController{}
		//指定controller
		userGroup.GET("/query",user.UserList)
		userGroup.POST("/register",user.Register)
		userGroup.GET("/proto",user.ProtoDemo)
		userGroup.POST("/parsing_proto",user.ParsingProtoDemo)

	}
}
