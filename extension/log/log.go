/**
@Time : 2019/10/14 11:32
@Author : liwei
@File : Logger
@Software: GoLand
*/

package log

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

// beego 日志配置结构体
type LoggerConfig struct {
	FileName            string `json:"filename"` //将日志保存到的文件名及路径
	Level               int    `json:"level"`    // 日志保存的时候的级别，默认是 Trace 级别
	Maxlines            int    `json:"maxlines"` // 每个文件保存的最大行数，若文件超过maxlines，则将日志保存到下个文件中，为0表示不设置。默认值 1000000
	Maxsize             int    `json:"maxsize"`  // 每个文件保存的最大尺寸，若文件超过maxsize，则将日志保存到下个文件中，为0表示不设置。默认值是 1 << 28, //256 MB
	Daily               bool   `json:"daily"`    // 设置日志是否每天分割一次，默认是 true
	Maxdays             int    `json:"maxdays"`  // 设置保存最近几天的日志文件，超过天数的日志文件被删除，为0表示不设置，默认保存 7 天
	Rotate              bool   `json:"rotate"`   // 是否开启 logrotate，默认是 true
	Perm                string `json:"perm"`     // 日志文件权限
	RotatePerm          string `json:"rotateperm"`
	EnableFuncCallDepth bool   `json:"-"` // 输出文件名和行号
	LogFuncCallDepth    int    `json:"-"` // 函数调用层级
}


func LogsInit() {
	var logCfg = LoggerConfig{
		FileName:            "logs/log.log",
		Level:               logs.LevelDebug,
		Daily:               true,
		EnableFuncCallDepth: true,
		LogFuncCallDepth:    3,
		RotatePerm:          "777",
		Perm:                "777",
		Maxdays:			30,
		Maxlines:           1000000000,


	}

	// 设置beego log库的配置
	b, _ := json.Marshal(&logCfg)
	_ = logs.SetLogger(logs.AdapterFile, string(b))
	//logs.Async() //为了提升性能, 可以设置异步输出
	logs.Async(1e3) //异步输出允许设置缓冲 chan 的大小
}

func Error(str string){
	logs.Error(str)
}



func EventDate()(string){
	//设置时区
	loc, _ := time.LoadLocation(beego.AppConfig.String("TIMEZONE"))
	tm :=time.Now().In(loc)
	return tm.Format(beego.AppConfig.String("FormatTable"))
}

func ReadLogs(c *gin.Context){
	var html string
	passWord:=c.DefaultQuery("auth_key","")
	c.Header("Content-Type", "text/html; charset=utf-8")
	if passWord=="" || passWord!=os.Getenv("LOG_PASSWORD"){
		html="<h1>error</h1>"
		c.String(http.StatusOK ,html)
		return
	}

	html +=`<html><head>`
	html +=`<link rel="stylesheet" type="text/css" href="/assets/css/H-ui.min.css" />`
	html +=`<link rel="stylesheet" type="text/css" href="/assets/css/H-ui.admin.css" />`
	html +=`</head>`
	html +=`<table  class="table table-border table-bg table-bordered" >
  			<thead>
    		<tr><th width="20%">日期</th><th>查看</th></tr>
  		</thead>
  		<tbody>`

	className :=make(map[int]string)
	className[0] = "active"
	className[-1] = "success"
	className[-2] = "warning"
	className[-3] = "danger"
	className[-4] = "active"
	className[-5] = "success"
	className[-6] = "warning"
	html +=`<tr class="`+className[-1]+`"><th>今天</th><td> <a href="/logs_info?date=log&auth_key=`+passWord+`">点击查看</a></td></tr>`
	for i:=0;i>=-6;i--{
		date :=time.Now().AddDate(0,0,i).Format("20060102")
		logFilePath :="cache/"
		logFileName :="log."+date
		fileName := path.Join(logFilePath, logFileName)
		_, err :=os.Stat(fileName)
		if err !=nil{
			logs.Error("不存在日志文件")
			continue
		}else{
			html +=`<tr class="`+className[i]+`"><th>`+date+`</th><td> <a href="/logs_info?date=`+date+`&auth_key=`+passWord+`">点击查看</a></td></tr>`
		}

	}

	html +=`</tbody></table>`
	html +=`</html>`

	c.String(http.StatusOK ,html)
	return
}


func ReadLogInfo(c *gin.Context){
	var html string
	passWord:=c.DefaultQuery("auth_key","")
	c.Header("Content-Type", "text/html; charset=utf-8")
	if passWord=="" || passWord!=os.Getenv("LOG_PASSWORD"){
		c.String(http.StatusOK ,html)
		return
	}

	html +=`<html><head>`
	html +=`<link rel="stylesheet" type="text/css" href="/assets/css/H-ui.min.css" />`
	html +=`<link rel="stylesheet" type="text/css" href="/assets/css/H-ui.admin.css" />`
	html +=`</head>`
	html +=`<table  class="table table-border table-bg table-bordered" >
  			<thead>
    		<tr><th width="20%">ENV</th><th>Actions</th></tr>
  		</thead>
  		<tbody>`

	date:=c.DefaultQuery("date",time.Now().Format("20060102"))
	lines := int64(200)
	logFilePath :="cache/"
	logFileName :="log."+date
	fileName := path.Join(logFilePath, logFileName)
	file,err:=os.Open(fileName)
	if err != nil {
		html = ""
		c.String(http.StatusOK,html)
		return
	}

	if err != nil {
		log.Println(err)
		return
	}

	fileInfo, _ := file.Stat()
	buf := bufio.NewReader(file)
	offset := fileInfo.Size() % 8192
	data := make([]byte, 8192) // 一行的数据
	totalByte := make([][][]byte, 0)
	readLines := int64(0)
	for i := int64(0); i <= fileInfo.Size() / 8192; i++{
		readByte := make([][]byte, 0) // 读取一页的数据
		file.Seek(fileInfo.Size() - offset - 8192*i, io.SeekStart)
		data = make([]byte, 8192)
		n, err := buf.Read(data)
		if err == io.EOF {
			if strings.TrimSpace(string(bytes.Trim(data, "\x00"))) != ""{
				readLines++
				readByte = append(readByte, data)
				totalByte = append(totalByte, readByte)
			}
			if readLines > lines{
				break
			}
			continue
		}
		if err != nil {
			log.Println("Read file error:", err)
			return
		}
		strs := strings.Split(string(data[:n]), "\n")
		if len(strs) == 1{
			b := bytes.Trim([]byte(strs[0]), "\x00")
			if len(b) == 0{
				continue
			}
		}
		if (readLines + int64(len(strs))) > lines{
			strs = strs[int64(len(strs))-lines+readLines:]
		}
		for j:=0;j<len(strs);j++{
			readByte = append(readByte, bytes.Trim([]byte(strs[j]+"\n"),"\x00"))
		}
		readByte[len(readByte)-1] = bytes.TrimSuffix(readByte[len(readByte)-1], []byte("\n"))
		totalByte = append(totalByte, readByte)
		readLines += int64(len(strs))

		if readLines >= lines{
			break
		}
	}
	totalByte = reverseByteArray(totalByte)
	str :=byteArrayToString(totalByte)
	for _, v := range str {
		if len(v)==0{
			continue
		}

		html +=`<tr class=""><td>local</td><td>`+v+`</td></tr>`
	}


	html +=`</tbody></table>`
	html +=`</html>`
	c.String(http.StatusOK,html)
	return
}


func reverseByteArray(s [][][]byte) [][][]byte {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
	return s
}

func byteArrayToString(buf [][][]byte) []string {
	str := make([]string, 0)
	for _, v := range buf {
		for _, vv := range v {
			str = append(str, string(vv))
		}
	}
	return str
}