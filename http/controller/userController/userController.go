package userController

import (
	"fmt"
	"gin.test/extension/log"
	"gin.test/http/controller"
	"gin.test/http/proto/httpReponseData"
	"gin.test/modules/user/userModel"
	"gin.test/modules/user/userRepositories"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"strconv"
)

type UserController struct {
	controller.GinBaseController
}

func (this *UserController) Add(ctx *gin.Context){
	this.LocalDate()
}

func (this *UserController) UserList(ctx *gin.Context)  {
	//获取所有url里的参数
	//items:=ctx.Request.URL.Query()
	//os.Exit(1)

	log.PError("xxx")
	id ,_:=  strconv.Atoi(ctx.Query("user_id"))

	//fmt.Println(id)
	//os.Exit(1)
	//name :=ctx.Query("name")
	//timeU :=this.LocalDate()
	data,err := userRepositories.GetUserById(id)
	if  err!=nil{
		ctx.JSON(200, controller.HTTP_ERROR_NOT_FOUND_USER)
		return
	}
	ctx.JSON(
		200,
		data,
	)
}





func (this *UserController) Register(c *gin.Context){
	//form 绑定验证
	var loginForm userModel.UserForm
	err:=c.ShouldBind(&loginForm)

	if err != nil {
		resp:= this.HTTP_ERROR_WITH_DATA(controller.HTTP_ERROR_REQUEST_HEADER_FAIL)
		c.JSON(200, resp)
		return
	}
	c.JSON(200, controller.HTTP_SUCCESS)
	return

}

func (this *UserController) ProtoDemo(c *gin.Context)  {
	//data :=map[string]string{"user_name":"天才","age":"18",}
	//result :=new(httpReponseData.User)
	//result.Name = "天才"
	//result.Age= 12

	data:=&httpReponseData.User{
		Name:"天才",
		Age:12,
	}
	c.ProtoBuf(http.StatusOK, data)
	//c.JSON(200, data)
	return

}

func (this *UserController)ParsingProtoDemo(c *gin.Context)  {
	resp, err := http.Get("http://127.0.0.1:1213/v1/user/proto")
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			user := &httpReponseData.User{}
			proto.UnmarshalMerge(body, user)
			fmt.Println(user.Age)
		}
	}
}