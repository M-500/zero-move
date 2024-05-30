package base

import (
	"sync"
	"testing"
)

// @Description
// @Author 代码小学生王木木

func TestReadMutex(t *testing.T) {
	lock := sync.Mutex{}
	lock.Lock()
}
