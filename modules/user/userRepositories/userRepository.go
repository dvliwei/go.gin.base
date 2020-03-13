//数据处理
package userRepositories

import (
	"errors"
	"fmt"
)

func GetUserList() map[string]string  {
	user :=map[string]string{}
	user["id"] = "1"
	user["name"] = "小明"

	return user
}

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
}
func GetUserById(id int) (User, error)  {
	a := make([]map[string]string, 1)
	a[0] = make(map[string]string)
	a[0]["title"] = "a"
	a[0]["desc"] = "hello"
	fmt.Printf("value is %+v\n", a)

	var users []User
	var data User
	data.ID = 1
	data.Name ="小明"
	users = append(users,data)
	data.ID = 2
	data.Name ="小亮"
	users = append(users,data)
	if len(users)<=id{
		err :=errors.New("not found user")
		return data,err
	}
	return users[id],nil
}