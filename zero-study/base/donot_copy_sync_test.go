package base

import (
	"sync"
	"testing"
)

// @Description
// @Author 代码小学生王木木

func TestCopySync(t *testing.T) {
	var lock sync.Mutex
	lock.Lock() // 先加锁
	b := lock
	b.Lock()
	b.Unlock()
}
