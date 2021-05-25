/**
 * @ClassName kernel
 * @Description //TODO 定时任务
 * @Author liwei
 * @Date 2020/6/15 16:17
 * @Version gin.test V1.0
 **/

package console

import (
	"gin.test/console/command"
	"gin.test/console/command/logCleanUpCommand"
	"github.com/astaxie/beego/toolbox"
)
var GameTK1 *toolbox.Task
var LogCleanUpTask *toolbox.Task

func init(){
	GameTK1 = toolbox.NewTask("tk1","* * * * * *", command.TestTaskCommand)
	LogCleanUpTask = toolbox.NewTask("清理日志","0 0 1 * * *", logCleanUpCommand.CleanUpLog)
}