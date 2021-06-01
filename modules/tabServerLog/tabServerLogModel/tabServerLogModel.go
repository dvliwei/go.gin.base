/**
 * @ClassName tabServerLogModel
 * @Description //TODO 
 * @Author liwei
 * @Date 2020/12/16 10:44
 * @Version game.cdkey.api V1.0
 **/

package tabServerLogModel

import (
	"go.translation.api/conf"
	"go.translation.api/extension/md5"
	"go.translation.api/extension/redis"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
	"strings"
	"time"
)

type TabServerLogModel struct {
	gorm.Model
	LogType  string `gorm:"column:log_type;type:varchar(32);index;comment:'日志类型'"`
	LogKey   string `gorm:"column:log_key;type:varchar(100);index;comment:'日志说明'"`
	LogValue string `gorm:"column:log_value;type:longText;comment:'日志内容'"`
}

func (TabServerLogModel)TableName()string  {
	return "tab_server_log"
}


//@Title CeckAndCreateTable
//@Description 检查和创建表
//@Description 同样的账号密码可以注册不同的游戏
//@param tableName string
func (TabServerLogModel) CheckAndCreateTable(tableName string){
	cacheKey:=os.Getenv("REDIS_PREFIX")+md5.StrMd5(tableName)
	if redis.HasRedis(cacheKey){
		return

	}else{
		if !conf.SERVER_LOG_DB.HasTable(tableName){
			conf.SERVER_LOG_DB.Table(tableName).Set("gorm:table_options", "ENGINE=Myisam DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='服务器日志'").
				CreateTable(&TabServerLogModel{})
		}else{
			//dbLog.ServerInfoMap("存储tab_user表是否存在判断结果", map[string]interface{}{"tabName":tableName,"key":cacheKey})
			redis.SetDataToRedisWithSecond(cacheKey,1,86400)
		}
	}

}

// @Description 插入数据
// @Description
// @param userId int
// @param model TabUser
// @return bool 成功后返回 新的信息 false
func(table *TabServerLogModel)NewData(model TabServerLogModel){
	loc, _ := time.LoadLocation(os.Getenv("TIMEZONE"))
	tm :=time.Now().In(loc)
	day :=tm.Format(os.Getenv("FORMATDATE"))
	tableName := "tab_server_log_"+day
	table.CheckAndCreateTable(tableName)
	//db.NewRecord(model) // => 成功后返回false
	//可以考虑异步固化数据到数据库
	//处理数据库链接错误
	conf.SERVER_LOG_DB.Table(tableName).Create(&model).NewRecord(model)
	return
}

/**
* @Title 清空日志表
* @Description:
* @Param:
* @return:
* @Author: liwei
* @Date: 2020/12/2
**/
func(table *TabServerLogModel)DropDate(logDate string){
	tableName := "tab_server_log_"+logDate
	table.CheckAndCreateTable(tableName)
	if conf.SERVER_LOG_DB.HasTable(tableName){
		conf.SERVER_LOG_DB.DropTableIfExists(tableName)
	}
	return
}

func(table TabServerLogModel)CreateTable(logDate string){
	tableName := "tab_server_log_"+logDate
	table.CheckAndCreateTable(tableName)
	return
}


func (table TabServerLogModel)GetLogTabName()[]string  {
	items:=make([]string,0)
	sql:="select table_name from information_schema.tables where table_schema='project_cdkey_shop' and table_name like '%tab_server_log_%' ORDER BY TABLE_NAME desc"
	rows,err:=conf.SERVER_LOG_DB.Raw(sql).Rows()
	defer  rows.Close()
	if err!=nil{
		return items
	}
	var table_name string
	for rows.Next(){
		rows.Scan(&table_name)
		table_name=strings.Replace(table_name,"tab_server_log_","",-1)
		items=append(items,table_name)
	}
	return items
}

func (table TabServerLogModel)GetDayLogsWithPage(logDate string,page int)([]TabServerLogModel, int)  {
	tableName := "tab_server_log_"+logDate
	table.CheckAndCreateTable(tableName)
	var models []TabServerLogModel
	count:=0
	db :=conf.SERVER_LOG_DB.Table(tableName)
	db.Count(&count)
	pageSize,_:= strconv.Atoi(os.Getenv("PAGE_SIZE"))
	db.Order("id desc").Offset((page-1)*pageSize).Limit(pageSize).Find(&models)
	return models,count
}