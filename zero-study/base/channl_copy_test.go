package base

import (
	"fmt"
	"testing"
	"time"
)

// @Description channel本质是值拷贝
// @Author 代码小学生王木木

type user struct {
	name string
	age  int
}

var u = user{name: "你大爷", age: 88}
var g = &u

func changeUserName(p *user) {
	fmt.Println("修改数据", u)
	p.name = "蔡徐坤"
}

func displayUser(uch <-chan *user) {
	time.Sleep(time.Second * 2)
	fmt.Println("收到了管道里的数据", <-uch)
}
func TestChannelCopy(t *testing.T) {

}
