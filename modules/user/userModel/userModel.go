//数据模型

package userModel

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)



type UserForm struct {
	Email string `form:"email" binding:"required,email"`  // 绑定表单类型的数据，且字段必须存在
	Password string `form:"password" binding:"required,min=10,max=32"`
}

type UserModel struct {
	gorm.Model
	Email string `gorm:"type:varchar(100);unique_index;comment:'邮箱'"`
	Password string `gorm:"type:varchar(64);"`
}

//ModelFieldTran 模型名称转换
type ModelFieldTran map[string]string

//FieldTrans 模型字段转换
func (u UserModel) FieldTrans() ModelFieldTran {
	m := ModelFieldTran{}
	m["Password"] = "用户密码"
	m["Email"] = "邮箱"
	return m
}

func Validate(req UserModel)error{

	validate := validator.New()

	err :=validate.Struct(req)
	fmt.Println("====")
	fmt.Println(err)
	if err!=nil{
		return err
	}
	return nil
}