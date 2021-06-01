/**
 * @ClassName logCleanUpCommand
 * @Description //TODO 
 * @Author liwei
 * @Date 2020/12/16 10:50
 * @Version game.cdkey.api V1.0
 **/

package logCleanUpCommand

import (
	"go.translation.api/extension/dbLog"
	"os"
	"time"
)

func CleanUpLog() error{
	loc, _ := time.LoadLocation(os.Getenv("TIMEZONE"))
	tm :=time.Now().AddDate(0,0,-7).In(loc)
	logDate :=tm.Format(os.Getenv("FORMATDATE"))
	dbLog.ServerInfo("清理服务器日志",logDate)
	dbLog.DropLogData(logDate)

	//创建下一日表
	tm2 :=time.Now().AddDate(0,0,1).In(loc)
	logDate2 :=tm2.Format(os.Getenv("FORMATDATE"))
	dbLog.ServerInfo("创建第二日的日志表",logDate2)
	dbLog.CreateLogTable(logDate2)

	return nil
}