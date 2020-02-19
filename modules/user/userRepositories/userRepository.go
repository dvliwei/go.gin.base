//数据处理
package userRepositories

import "fmt"

func GetUserList() map[string]string  {
	user :=map[string]string{}
	user["id"] = "1"
	user["name"] = "小明"

	return user
}

func GetUserById(id string) map[string]string  {
	a := make([]map[string]string, 1)
	a[0] = make(map[string]string)
	a[0]["title"] = "a"
	a[0]["desc"] = "hello"
	fmt.Printf("value is %+v\n", a)

	users := make([]map[string]string,1)
	users[0] = make(map[string]string)
	users[0]["id"] = "1"
	users[0]["name"]="小明"

	return users[0]
}