package model

import (
	"fmt"
	"os"
)

type User struct {
	Username string
	Password string
}

func Addusername(username string) {
	file, err := os.OpenFile("./用户名.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("文件错误,错误为:%v\n", err)
		return
	}
	defer file.Close()
	n, err := file.WriteString(username + "\n")
	if err != nil {
		fmt.Println(n)
		return
	}

}
func Addpassword(password string) {
	file, err := os.OpenFile("./密码.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("文件错误,错误为:%v\n", err)
		return
	}
	defer file.Close()
	n, err := file.WriteString(password + "\n")
	if err != nil {
		fmt.Println(n)
		return
	}

}
