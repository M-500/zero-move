package base

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// @Description
// @Author 代码小学生王木木

func TestContex(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ctx.Done()
}

func TestContextNo(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	go Run(ctx)
	time.Sleep(time.Second * 30) // 等待30秒
}
func Run(ctx context.Context) {
	for true {
		select {
		case <-ctx.Done():
			fmt.Println("执行结束")
			return
		default:
			fmt.Println("芽儿哟！")
			time.Sleep(time.Second)
		}
	}

}

func TestContextCancel(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())

	ctx2 := context.WithValue(ctx, "name", "叼毛")
	go RunV1(ctx2)
	time.Sleep(time.Second * 5)
	cancelFunc() // 5秒后调用取消 结束运行
}

func RunV1(ctx context.Context) {
	s, ok := ctx.Value("name").(string)
	if !ok {
		fmt.Println("类型转换错误")
	}
	fmt.Println(s)
}
