/**
 * @ClassName httpMiddleware
 * @Description  http 请求 header 头信息验证 //TODO
 * @Author liwei
 * @Date 2020/3/13 15:44
 * @Version go.translation.api V1.0
 **/

package middleware

import (
	"go.translation.api/http/controller"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config/env"
	"github.com/gin-gonic/gin"
)

/**
* @Title HttpHeaderVerification
* @Description:   请求头验证 gin使用中间件出错后不能用return终止，而应该使用Abort实现
* @Param:  
* @return:  
* @Author: liwei
* @Date: 2020/3/13 
**/
func HttpHeaderVerification()  gin.HandlerFunc {
	return func(context *gin.Context) {
		httpToken := context.Request.Header.Get("http-token")
		if httpToken!=env.Get("http-token","140d7a33b5f31259d4d035dd3fb34b9118daf551"){
			allowCross(context)
		}
		context.Next()
	}
}

//错误返回
func allowCross(ctx *gin.Context) {
	ctx.Request.Header.Set("Cache-Control", "no-store")
	ctx.Request.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Request.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE,OPTIONS")
	ctx.Request.Header.Set("Access-Control-Allow-Headers", "Authorization")
	ctx.Request.Header.Set("WWW-Authenticate", `Bearer realm="`+beego.AppConfig.String("HostName")+`" error="Authorization" error_description="invalid Authorization"`)
	ctx.JSON(200,controller.HTTP_ERROR_REQUEST_HEADER_FAIL)
	ctx.Abort()
}