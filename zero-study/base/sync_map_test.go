package base

import (
	"sync"
	"testing"
)

// @Description
// @Author 代码小学生王木木

func TestSyncMap(t *testing.T) {
	s := sync.Map{}
	s.Load("无鳞鳞")
}
