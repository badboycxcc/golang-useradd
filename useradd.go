package main

import (
	"flag"
	"fmt"

	wapi "github.com/iamacarpet/go-win64api"
)

func main() {

	// 用户
	var username string
	// 密码
	var password string
	// 管理组
	var groupname string

	flag.StringVar(&username, "u", "admin123", "用户名,默认 admin123")
	flag.StringVar(&password, "p", "admin123", "密码,默认 admin123")
	flag.StringVar(&groupname, "g", "administrators", "用户组,默认 administrators")
	flag.Parse()

	fullname := ""
	groupuser := []string{username}
	wapi.UserAdd(username, fullname, password)
	wapi.LocalGroupAddMembers(groupname, groupuser)

	// 打印
	fmt.Printf("Add User username=%v password=%v group=%v\n", username, password, groupname)

}
