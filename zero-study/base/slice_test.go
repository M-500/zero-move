package base

import (
	"fmt"
	"reflect"
	"testing"
)

// @Description
// @Author 代码小学生王木木

func add(arr []int) {
	arr[0] = 2
}

func TestSlice(t *testing.T) {
	arr := []int{1, 2, 3}
	add(arr)
	t.Log(arr, reflect.TypeOf(arr).Kind())
}

func TestDemo(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]
	s2 = append(s2, 100)
	s2 = append(s2, 200)
	s1[2] = 20
	t.Log(s1)
	t.Log(s2)
	t.Log(slice)
}

func TestDemo2(t *testing.T) {
	s := make([]int, 0, 2)
	doSomething1(s)
	fmt.Println(s)
}
func doSomething1(a []int) {
	a = append(a, 1)
}
