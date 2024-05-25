package base

import (
	"reflect"
	"testing"
)

// @Description
// @Author 代码小学生王木木

func add(arr [3]int) {
	arr[0] = 2
}

func TestSlice(t *testing.T) {
	arr := [3]int{1, 2, 3}
	add(arr)
	t.Log(arr, reflect.TypeOf(arr).Kind())
}
