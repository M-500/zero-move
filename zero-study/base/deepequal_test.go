package base

import (
	"reflect"
	"testing"
)

// @Description
// @Author 代码小学生王木木

func TestEqual(t *testing.T) {
	reflect.DeepEqual(1, 2)
}

func TestArrayEqual(t *testing.T) {
	//a := make(map[[1]int]string, 0)
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	t.Log(a == b)
	//a1 := []int{1, 2, 3}
	//a2 := []int{1, 2, 3}
	//t.Log(a1 == a2)
}

//type Dog struct {
//	Name string
//	Age  int
//}
//
//func TestStructEqualV1(t *testing.T) {
//	a := User{Name: "小黑子", Age: 25}
//	b := Dog{Name: "小黑子", Age: 25}
//	t.Log(a == b)
//}

type User struct {
	Name      string
	Age       int
	PhoneList []string
	addr      string
}

func TestStructEqual(t *testing.T) {
	a := User{Name: "小黑子", Age: 25, PhoneList: []string{"13533333333"}, addr: "1"}
	b := User{Name: "小黑子", Age: 25, PhoneList: []string{"13533333333"}, addr: "2"}
	//t.Log(a == b) // 无法使用 == 进行比较
	equal := reflect.DeepEqual(a, b)
	t.Log("是否相等？", equal)
}
