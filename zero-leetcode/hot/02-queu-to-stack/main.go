package _2_queu_to_stack

// @Description 请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。
// https://leetcode.cn/problems/implement-stack-using-queues/description/
// @Author 代码小学生王木木
type MyStack struct {
	master []int
	slave  []int
}

func Constructor() MyStack {
	return MyStack{
		master: make([]int, 0),
		slave:  make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	if len(this.master) < 1 {
		this.master = append(this.master, x)
		return
	}
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return -1
	}
	return 0
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return -1
	}
	return 0
}

func (this *MyStack) Empty() bool {
	if len(this.slave) == 0 && len(this.master) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
