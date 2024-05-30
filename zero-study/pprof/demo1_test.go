package pprof

import (
	"log"
	"os"
	"runtime/pprof"
	"sync"
	"testing"
)

// @Description
// @Author 代码小学生王木木

type UserModel struct {
	Name string
	Age  int
}

func TestPProf(t *testing.T) {
	//采样cpu运行状态
	f, err := os.Create("cpu.pprof")
	if err != nil {
		log.Fatal("无法创建 CPU profile: ", err)
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal("采集CPU信息失败: ", err)
	} // 采样CPU
	defer pprof.StopCPUProfile()

	f2, err2 := os.Create("mem.pprof")
	if err2 != nil {
		log.Fatal("无法创建 Memory profile: ", err)
	}
	err = pprof.WriteHeapProfile(f2)
	if err != nil {
		log.Fatal("采集内存信息失败: ", err)
	} // 采样内存

	// 创建一个channel 100个容量
	ch := make(chan UserModel, 100)
	var wg sync.WaitGroup
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		u := UserModel{
			Name: "",
			Age:  0,
		}
		go func(u UserModel) {
			defer func() {
				wg.Done()
			}()
			ch <- u
		}(u)
	}
	wg.Wait()
}
