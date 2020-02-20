package controller

import (
	"time"
)

//NewTodoController - create todo controller with mehtod dealing with todo item

func BindingController() *GinBaseController{

	return &GinBaseController{}
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