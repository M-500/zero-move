package _1_stack_to_queu

// @Description https://leetcode.cn/problems/implement-queue-using-stacks/  用两个栈实现一个队列
// @Author 代码小学生王木木

type MyQueue struct {
	addStack []int
	delStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		addStack: make([]int, 0),
		delStack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.addStack = append(this.addStack, x)
}

// Pop
//
//	@Description: 移除队头元素并返回
//	@receiver this
//	@return int
func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}
	if len(this.delStack) == 0 {
		this.moveToDelStack() // 将数据栈的数据全部弹出，压入删除栈
	}
	i := this.delStack[len(this.delStack)-1]
	this.delStack = this.delStack[:len(this.delStack)-1]
	return i
}

func (this *MyQueue) moveToDelStack() {
	if this.Empty() {
		return
	}
	if len(this.addStack) == 0 {
		return
	}
	for len(this.addStack) > 0 {
		tmp := this.addStack[len(this.addStack)-1]
		this.addStack = this.addStack[:len(this.addStack)-1]
		this.delStack = append(this.delStack, tmp)
	}
}

// Peek
//
//	@Description: 返回队列开头的元素 不会移除这个元素，仅仅获取并且打印
//	@receiver this
//	@return int
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}
	if len(this.delStack) == 0 {
		this.moveToDelStack() // 将数据栈的数据全部弹出，压入删除栈
	}
	i := this.delStack[len(this.delStack)-1]
	return i
}

// Empty
//
//	@Description: 如果队列为空，返回 true ；否则，返回 false
//	@receiver this
//	@return bool
func (this *MyQueue) Empty() bool {
	if len(this.delStack) == 0 && len(this.addStack) == 0 {
		return true
	}
	return false
}
