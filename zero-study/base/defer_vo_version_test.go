package base

import "testing"

// @Description
// @Author 代码小学生王木木

// 分别在go1.21 1.22 下运行
func TestDefer1_21(t *testing.T) {
	// 在 1.21 之下，输出十个 10
	// 在 1.22 之下，输出 9 - 0
	// 1.22 的时候，循环变量 j 每次都是新的（你可以理解为，每次迭代都用一块新的内存）
	for i := 0; i < 10; i++ {
		defer func() {
			t.Log(i)
		}()
	}
}
