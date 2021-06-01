/**
 * @ClassName logController
 * @Description //TODO 
 * @Author liwei
 * @Date 2021/6/1 10:59
 * @Version go.translation.api V1.0
 **/

package log_controller

import (
	"github.com/gin-gonic/gin"
	"go.translation.api/http/admin/utils"
	"go.translation.api/http/controller"
	"go.translation.api/modules/tabServerLog/tabServerLogModel"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"
)

type LogController struct {
	controller.GinBaseController
}


func(this *LogController) LogList(ctx *gin.Context){

	var serverLog tabServerLogModel.TabServerLogModel
	items:= serverLog.GetLogTabName()
	loc, _ := time.LoadLocation(os.Getenv("TIMEZONE"))
	tm2 :=time.Now().In(loc)
	logDate :=tm2.Format(os.Getenv("FORMATDATE"))
	page:=this.GetInt("page",0,ctx)
	date:=this.GetString("date",logDate,ctx)
	dayLogs,count :=serverLog.GetDayLogsWithPage(date,page)
	pageSize,_ :=strconv.Atoi(os.Getenv("PAGE_SIZE"))
	pagination := utils.NewPagination(ctx.Request, count,pageSize)
	seachValue:=this.GetString("keyword","",ctx)
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
		"items":items,
		"day_logs":dayLogs,
		"all_count":count,
		"day":date,
		"seach_value":seachValue,
		"pages":template.HTML(pagination.Pages()),
	})
}