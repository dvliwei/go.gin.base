/**
 * @ClassName redis
 * @Description //TODO 
 * @Author liwei
 * @Date 2020/4/7 10:38
 * @Version gin.test V1.0
 **/

package redis

import (
	"gin.test/conf"
	"gin.test/extension/log"
	"time"
)

/**
* @Title GetRedisDataByKey
* @Description:  通过key查询redis
* @Param:
* @return:
* @Author: liwei
* @Date: 2020/4/7
**/
func GetRedisDataByKey(key string)(string, error){
	str ,err:= conf.PREDIS.Get(key).Result()
	return str,err
}

/**
* @Title SetDataToRedisWithKey
* @Description:  保存数据到redis 存储时间按照毫秒结算
* @Param:
* @return:
* @Author: liwei
* @Date: 2020/4/7
**/
func SetDataToRedisWithMillisecond(key string,value interface{},timeOut time.Duration){
	err:=conf.PREDIS.SetNX(key ,value,time.Millisecond*timeOut).Err()
	if err!=nil{
		log.PError("set redis error"+err.Error())
	}
}

/**
* @Title SetDataToRedisWithSecond
* @Description:   保存数据到redis 存储时间按照秒结算
* @Param:
* @return:
* @Author: liwei
* @Date: 2020/4/7
**/
func SetDataToRedisWithSecond(key string,value interface{},timeOut time.Duration){
	err:=conf.PREDIS.SetNX(key ,value,time.Second*timeOut).Err()
	if err!=nil{
		log.PError("set redis error"+err.Error())
	}
}

/**
* @Title DelRedisDataByKey
* @Description:  删除缓存
* @Param:
* @return:
* @Author: liwei
* @Date: 2020/4/7
**/
func DelRedisDataByKey(key string){
	err:=conf.PREDIS.Del(key).Err()
	if err!=nil{
		log.PError("del redis data error "+err.Error())
	}
}

/**
* @Title HasRedis
* @Description: 判断是否存在redis  Exists 存在1否则0
* @Param:
* @return:  bool
* @Author: liwei
* @Date: 2020/4/7
**/
func HasRedis(key string) bool{
	status,err:=conf.PREDIS.Exists(key).Result()
	if err!=nil{
		log.PError(" exists  redis error "+err.Error())
	}
	if status ==0{
		return false
	}
	return true
}