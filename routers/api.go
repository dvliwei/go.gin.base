package routers

import (
	"gin.test/extension/server"
	"gin.test/http/controller/userController"
	"gin.test/http/middleware"
	"github.com/gin-gonic/gin"
)


func Routers(){

	v1:= server.Server.Group("v1")
	//路由组中添加中间件
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


	}
}
