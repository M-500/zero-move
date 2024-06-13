package _3_last_n_linklist

// @Description
// @Author 代码小学生王木木

type LinkList struct {
	Val  int
	Next *LinkList
}

func GetLastN(head *LinkList, n int) *LinkList {
	if head == nil {
		return nil
	}
	left := head
	right := head
	// right先往后走N步
	for i := 0; i < n; i++ {
		right = right.Next
	}
	// left right一起往后走 right走到头了说明 left就是第N个节点
	for right != nil {
		left = left.Next
		right = right.Next
	}
	return left
}
