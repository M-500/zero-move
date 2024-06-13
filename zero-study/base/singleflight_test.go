//@Author: wulinlin
//@Description:
//@File:  singleflight_test
//@Version: 1.0.0
//@Date: 2024/06/11 19:09

package base

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

func TestSingleFight(t *testing.T) {
	var wg sync.WaitGroup
	var ThreadNum = 10

	wg.Add(ThreadNum)

	for i := 0; i < ThreadNum; i++ {
		go func() {
			defer wg.Done()
			data, err := load("key")
			if err != nil {
				t.Log(err)
			}
			t.Log(data)
		}()
	}
}

func load(key string) (any, error) {

}

func loadFromCache(key string) (any, error) {
	return "", errors.New("缓存没有数据，未命中")
}

func setCache(key string) {

}

func loadFromDB(key string) (any, error) {
	fmt.Println("查询数据库")
	return "数据库中有数据", nil
}
