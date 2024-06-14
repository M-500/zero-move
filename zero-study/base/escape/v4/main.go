package main

import "fmt"

type Love interface {
	SayLove()
}
type guanDong struct{}

func (f1 guanDong) SayLove() {
	fmt.Println("我吼中医类啊")
}

func main() {
	var guanXiChen Love
	guanXiChen = guanDong{}
	guanXiChen.SayLove() // 调用方法时，guanXiChen发生逃逸，因为方法是动态分配的
}
