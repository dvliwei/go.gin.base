package conf

import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

var PORT string //服务器端口

var SERVER_DB  *gorm.DB //数据库

var PREDIS *redis.Client //redis

var TIMEZONE string //时区

var SERVER_DATE string //服务器当前日期

var SERVER_DATE_TIME string //服务器当前时间
func init(){
	err:=godotenv.Load()
	if err!=nil{
		log.Fatal("Error loding .env file")
	}

	//服务器端口配置
	PORT =os.Getenv("HTTP_PORT")

	//服务器当前日期
	loc, _ := time.LoadLocation(os.Getenv("TIMEZONE"))
	tm :=time.Now().In(loc)
	SERVER_DATE = tm.Format("20060102")
	SERVER_DATE_TIME = tm.Format("2006-01-02 15:04:05")

	//数据库配置
	SERVER_DB,err =gorm.Open(os.Getenv("DB_CONNECTION"),
		os.Getenv("DB_USERNAME")+":"+
			os.Getenv("DB_PASSWORD")+"" +
			"@/"+os.Getenv("DB_DATABASE")+"?charset="+
			os.Getenv("DB_CHARSET")+"&parseTime=True&loc=Local")
	if err!=nil{
		log.Fatal("connection mysql fail", err.Error())
	}

	//redis配置
	PREDIS = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})


}
