/**
 * @ClassName dbLog
 * @Description //TODO 
 * @Author liwei
 * @Date 2020/12/16 10:53
 * @Version game.cdkey.api V1.0
 **/

package dbLog

import (
	"encoding/json"
	"gin.test/modules/tabServerLog/tabServerLogModel"
	"reflect"
	"strconv"
)

/**
* @Title 服务器正常信息
* @Description:
* @Param:
* @return:
* @Author: liwei
* @Date: 2020/12/2
**/
func ServerInfo(keyName string,logValue ...interface{}){
	var model tabServerLogModel.TabServerLogModel
	model.LogType = "I"
	model.LogKey = keyName
	var msg string
	for _,value :=range logValue{
		msg +=FormatMessage(value)
	}
	model.LogValue = msg

	go func() {
		model.NewData(model)
	}()
}

/**
* @Title 服务器错误信息
* @Description:
* @Param:
* @return:
* @Author: liwei
* @Date: 2020/12/2
**/
func ServerError(keyName string,logValue ...interface{}){
	var model tabServerLogModel.TabServerLogModel
	model.LogType = "E"
	model.LogKey = keyName
	var msg string
	for _,value :=range logValue{
		msg +=FormatMessage(value)
	}
	model.LogValue = msg
	go func() {
		model.NewData(model)
	}()
}




/**
* @Title ServerInfoMap
* @Description:json存储日志
* @Param:
* @return:
* @Author: liwei
* @Date: 2021/4/14
**/
func ServerInfoMap(keyName string,logMapValue map[string]interface{}){
	var model tabServerLogModel.TabServerLogModel
	model.LogType = "I"
	model.LogKey = keyName
	bValue,err:=json.Marshal(logMapValue)
	if err!=nil{
		model.LogValue = err.Error()
	}else{
		model.LogValue = string(bValue)
	}
	go func(model tabServerLogModel.TabServerLogModel) {
		model.NewData(model)
	}(model)
}

/**
* @Title ServerErrorMap
* @Description:  json存储错误日志
* @Param:
* @return:
* @Author: liwei
* @Date: 2021/4/14
**/
func ServerErrorMap(keyName string,logMapValue map[string]interface{}){
	var model tabServerLogModel.TabServerLogModel
	model.LogType = "E"
	model.LogKey = keyName
	bValue,err:=json.Marshal(logMapValue)
	if err!=nil{
		model.LogValue = err.Error()
	}else{
		model.LogValue = string(bValue)
	}
	go func(tabServerLogModel.TabServerLogModel) {
		model.NewData(model)
	}(model)
}

func DropLogData(logDate string)  {
	var model tabServerLogModel.TabServerLogModel
	model.DropDate(logDate)
}

func CreateLogTable(logDate string)  {
	var model tabServerLogModel.TabServerLogModel
	model.CreateTable(logDate)
}


func FormatMessage(a interface{})  (result string) {

	rt := reflect.TypeOf(a)

	switch rt.Kind() {
	case reflect.Slice:
		str,_:=json.Marshal(a)
		result =  string(str)
	case reflect.Array:
		str,_:=json.Marshal(a)
		result = string(str)
	case reflect.Map:
		str,_:=json.Marshal(a)
		result = string(str)
	case reflect.String:
		result =a.(string)
	case reflect.Int:
		result = strconv.Itoa(a.(int))
	case reflect.Bool:
		result =strconv.FormatBool(a.(bool))
	default:
		result = a.(string)
	}
	return result
}
