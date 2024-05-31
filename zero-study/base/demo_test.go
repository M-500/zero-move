package base

import "testing"

// @Description
// @Author 代码小学生王木木
const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

func TestDemo1(t *testing.T) {

	t.Log("草了mutexLocked", mutexLocked)
	t.Log("草了mutexWoken", mutexWoken)
	t.Log("草了mutexStarving", mutexStarving)
	t.Log("草了mutexWaiterShift", mutexWaiterShift)

}
