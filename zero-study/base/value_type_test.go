package base

import (
	"fmt"
	"testing"
)

// @Description 值类型和引用类型
// @Author 代码小学生王木木

func TestValueType(t *testing.T) {
	a := 10
	b := a // 这里b是a的拷贝，a和b是独立的
	b = 20
	fmt.Println(a) // 输出10
	fmt.Println(b) // 输出20
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1 // 这里arr2是arr1的拷贝，arr1和arr2是独立的
	arr2[0] = 100
	fmt.Println(arr1) // 输出[1 2 3]
	fmt.Println(arr2) // 输出[100 2 3]
}
