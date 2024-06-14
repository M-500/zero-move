package main

import (
	"fmt"
	"runtime"
	"time"
)

// @Description
// @Author 代码小学生王木木

func longRunningTask() {
	for i := 0; i < 1000000000; i++ {
		// 每隔一定次数主动让出处理器
		//if i%1000000 == 0 {
		//	runtime.Gosched()
		//}
		fmt.Println("哈哈哈哈")
		time.Sleep(time.Second)
	}
	fmt.Println("长任务运行结束")
}

func main() {
	// 设置全局只有一个M
	runtime.GOMAXPROCS(1)
	go longRunningTask()
	go func() {
		for {
			fmt.Println("我可以执行一下？")
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 10)
}
