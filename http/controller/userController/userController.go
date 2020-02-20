package userController

import (
	"fmt"
	"gin.test/http/controller"
	"gin.test/modules/user/userModel"
	"gin.test/modules/user/userRepositories"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	controller.GinBaseController
}

func (this *UserController) Add(ctx *gin.Context){
	this.LocalDate()
}

func (this *UserController) UserList(ctx *gin.Context)  {
	id:= ctx.Query("id")


	fmt.Println(id)
	//os.Exit(1)
	//name :=ctx.Query("name")
	//timeU :=this.LocalDate()
	data := userRepositories.GetUserById(id)
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