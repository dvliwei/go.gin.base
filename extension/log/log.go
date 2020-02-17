package log

import (
	"fmt"
	"gin.test/conf"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)



func  PInfo(str string){
	datetime :=  conf.SERVER_DATE_TIME
	log := saveLogInFile()
	log.Infof("%s	%13v	%s","[info]",datetime,str)
}


func  PError(str string){
	datetime :=  conf.SERVER_DATE_TIME
	log := saveLogInFile()
	log.Errorf("%s	%13v	%s","[error]",datetime,str)
}

func PDebug(str string){
	datetime :=  conf.SERVER_DATE_TIME
	log := saveLogInFile()
	log.Debugf("%s	%13v	%s","[error]",datetime,str)
}

func PWarn(str string)  {
	datetime :=  conf.SERVER_DATE_TIME
	log := saveLogInFile()
	log.Warnf("%s	%13v	%s","[error]",datetime,str)
}



func saveLogInFile() *logrus.Logger{
	logFilePath :="cache/"
	logFileName :="log."+conf.SERVER_DATE
	//日志文件

	fileName := path.Join(logFilePath, logFileName)
	os.Create(fileName)
	//写入文件
	LoggerFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	//实例话
	logger :=logrus.New()
	//设置输出
	logger.Out = LoggerFile
	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{})
	return logger
}