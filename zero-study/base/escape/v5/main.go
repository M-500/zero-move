package main

import (
	"fmt"
	"gorm.io/gorm"
)

func f() {
	x := 1
	go func() {
		println(x)
	}()
}

var db *gorm.DB

type UserModel struct {
	Name string
	Age  int
}

func hh(model *UserModel) UserModel {
	model.Name = "傻逼"
	return *model
}
func main() {
	u := UserModel{Age: 12}
	hh(&u)
	fmt.Println(u.Name)
}
