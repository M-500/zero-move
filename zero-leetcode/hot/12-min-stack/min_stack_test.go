package _2_min_stack

// @Description 实现最小栈的功能
// @Author 代码小学生王木木
type MinStack struct {
	datStack []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		datStack: []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.datStack = append(this.datStack, x)
	if len(this.minStack) == 0 || this.Min() > x {
		this.minStack = append(this.minStack, x)
	} else {
		this.minStack = append(this.minStack, this.Min())
	}
}

func (this *MinStack) Pop() {
	if len(this.datStack) > 0 {
		this.datStack = this.datStack[:len(this.datStack)-1]
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.datStack) > 0 {
		return this.datStack[len(this.datStack)-1]
	}
	return -1
}

func (this *MinStack) Min() int {
	if len(this.minStack) == 0 {
		return -1
	}
	return this.minStack[len(this.minStack)-1]
}
