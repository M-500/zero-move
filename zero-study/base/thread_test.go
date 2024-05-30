package base

import (
	"fmt"
	"testing"
	"time"
)

// @Description
// @Author 代码小学生王木木

// 一个协程panic 其他协程全部遭殃
func TestPaincThread(t *testing.T) {
	go func() {
		for true {
			fmt.Println("A协程正在运行")
		}
	}()

	go func() {
		time.Sleep(time.Second)
		panic("B协程噶了")
	}()
	time.Sleep(time.Second * 5)
}

func TestRecover(t *testing.T) {
	go func() {
		panic("直接报错")
	}()
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("recover兜底抓住了panic")
		}
	}()
}

func TestNice(t *testing.T) {
	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Println("协程出现异常了，我抓住了！")
			}
		}()
		panic("直接panic，不讲道理")
	}()
	fmt.Println("执行结束")
}
