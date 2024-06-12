package base

import (
	"sync"
	"testing"
)

// @Description
// @Author 代码小学生王木木

type MyMap[T comparable] struct {
	lock *sync.RWMutex
	m    map[T]any
}

func NewMyMap[T comparable](l int) *MyMap[T] {
	return &MyMap[T]{
		m:    make(map[T]any, l),
		lock: &sync.RWMutex{},
	}
}

func TestOldMap(t *testing.T) {
	var oldMap = make(map[int]int, 0)
	// 开启协程 死循环读map
	go func() {
		for {
			_ = oldMap[1]
		}
	}()
	// 开启协程，死循环写map
	go func() {
		for {
			oldMap[2] = 3 // 即使是不同的key
		}
	}()

	select {} // hold 主main协程，不让他退出
}

func TestUseSyncMap(t *testing.T) {

}
