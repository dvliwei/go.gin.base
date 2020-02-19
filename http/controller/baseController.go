package controller

import (
	"time"
)

//NewTodoController - create todo controller with mehtod dealing with todo item

func BindingController() *GinBaseController{

	return &GinBaseController{}
}

type GinBaseController struct {
}

func (this *GinBaseController) LocalDate() int64{
	return time.Now().Unix()
}