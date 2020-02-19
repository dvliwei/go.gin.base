package userController

import (
	"fmt"
	"gin.test/http/controller"
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