package controller

import (
	"github.com/gin-gonic/gin"
	"go.translation.api/extension/dbLog"
	"strconv"
	"time"
)

//NewTodoController - create todo controller with mehtod dealing with todo item

func BindingController() *GinBaseController{

	return &GinBaseController{}
}

type Foo struct {
	Name string

}

type GinBaseController struct {
}



type GinResult struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type GinResultNoData struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type GinResultErrorNoData  struct {
	Code int `json:"code"`
	Message string `json:"message"`
}


type GinResultError struct {
	Code int `json:"code"`
	Message string `json:"message"`
	OsMessage interface{} `json:"osMessage,omitempty"`
}


var HTTP_SUCCESS  = GinResultNoData{Code:1,Message:"ok"}
func (this *GinBaseController) HTTP_SUCCESS_WITH_DATA(data interface{})(GinResult){
	result := GinResult{}
	result.Code = 1
	result.Message = "ok"
	result.Data = data
	return result
}

//@Title HTTP_ERROR_THIRD_VERIFYID_WITH_DATA
//@Description 请求参数验证返回值
//@Param data interface{}
//@Return json
func (this *GinBaseController)HTTP_ERROR_WITH_DATA(data interface{}) (GinResultError) {
	result := GinResultError{}
	result.Code = -1
	result.Message = "err"
	result.OsMessage = data
	return result
}


func (this *GinBaseController) LocalDate() int64{
	return time.Now().Unix()
}


//参数错误10000开头
var HTTP_ERROR_REQUEST_HEADER_FAIL = GinResultError{Code:401,Message:"Unauthorized"}


//参数错误20000开头业务错误

var HTTP_ERROR_NOT_FOUND_USER = GinResultError{Code:20000,Message:"not  found user"}

var HTTP_ERROR_TRANSLATION_FAIL = GinResultError{Code:20001,Message:"translation  is fail"}


/**
* @Title GetString
* @Description:  获取 字符串参数
* @Param:
* @return:
* @Author: liwei
* @Date: 2020/5/29
**/
func (this *GinBaseController)GetString(key string,defaultvalue string,c *gin.Context) string{
	param:= c.Query(key)
	if param=="" {
		param = c.PostForm(key)
		if param==""{
			param = defaultvalue
		}
	}
	return param

}

/**
* @Title GetInt
* @Description:  获取int参数
* @Param:
* @return:
* @Author: liwei
* @Date: 2020/5/29
**/
func (this *GinBaseController)GetInt(key string,defaultvalue int,c *gin.Context)(int){
	param:= c.Query(key)
	if param=="" {
		param = c.PostForm(key)
		if param==""{
			return defaultvalue
		}
	}
	intValue,err:=strconv.Atoi(param)
	if err!=nil{
		dbLog.ServerError("非法的请求参数 预期值为int 实际上传参数为",param)
		return defaultvalue
	}
	return intValue
}

/**
* @Title GetFloat64ToInt
* @Description:  float64转int
* @Param:
* @return:
* @Author: liwei
* @Date: 2020/5/29
**/
func (this *GinBaseController) GetFloat64(key string,defaultvalue float64,c *gin.Context) float64  {
	param:= c.Query(key)
	if param=="" {
		param = c.PostForm(key)
		if param==""{
			return defaultvalue
		}
	}
	float64Value,err :=strconv.ParseFloat(param,32/64)
	if err!=nil{
		return defaultvalue
	}else{
		return float64Value
	}
}