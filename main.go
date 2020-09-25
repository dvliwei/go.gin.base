package main

import (
	"context"
	"fmt"
	"gin.test/conf"
	"gin.test/console"
	"gin.test/extension/log"
	"gin.test/extension/server"
	"gin.test/routers"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)



func main(){

	log.LogsInit()
	timeAt:=time.Now().UnixNano()
	fmt.Println(timeAt)
	go func() {
		fmt.Println("==============")
		fmt.Println(timeAt)
		fmt.Println("==============")
		fmt.Println(time.Now().UnixNano())
	}()


	//设置定时任务
	if os.Getenv("OPEN_TASK")=="open"{
		toolbox.AddTask("tk1",console.GameTK1)
		toolbox.StartTask()
		defer toolbox.StopTask()
	}



	//设置运行模式
	//gin.ReleaseMode  release
	//gin.DebugMode dug
	gin.SetMode(os.Getenv("GIN_MODE"))
	//初始化路由
	routers.Routers()

	//Go版本是1.8，你可能不需要使用这个库，考虑使用http.Server内置的Shutdown()方法进行优雅关闭
	srv:=&http.Server{
		Addr:":"+conf.PORT,
		Handler:server.Server,
	}
	logs.Info(" listen tcp :"+srv.Addr)

	go func() {
		//服务器链接
		if err:=srv.ListenAndServe();err!=nil && err!=http.ErrServerClosed{
			logs.Error("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logs.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logs.Error("Server Shutdown:", err)
	}

	logs.Info("Server exiting")
	//server.Server.Run(":"+conf.PORT)
}