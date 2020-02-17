package server

import (
	"github.com/gin-gonic/gin"
)


//配置http服务器
//中间件
//其中 Logger 是对日志进行记录，而 Recovery 是对有 painc时, 进行 500 的错误处理
//查看了源码之后，那么我们也就知道如何使用中间件了。
var Server *gin.Engine
func init(){
	Server = gin.New()
	//全局中间件
	//Server.Use(gin.Logger())
	// 使用 Recovery 中间件
	//Server.Use(gin.Recovery())

}