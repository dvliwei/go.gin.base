package main

import (
	"gin.test/conf"
	"gin.test/extension/server"
	"gin.test/routers"
)

func main(){


	//初始化路由
	routers.Routers()
	server.Server.Run(":"+conf.PORT)
}