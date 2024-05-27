package main

import "fmt"

// @Description
// @Author 代码小学生王木木

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var slow *ListNode
	cur := head
	for cur != nil {
		fast := cur.Next // 记录下一个指针
		cur.Next = slow  // 破坏掉规则
		slow = cur       // 慢指针前移
		cur = fast       // 下一步
	}
	return slow
}

func reverseListV1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var stack []*ListNode
	tag := head
	for tag != nil {
		stack = append(stack, tag)
		s := tag.Next
		tag.Next = nil
		tag = s
	}
	var ans = &ListNode{}
	for i := len(stack) - 1; i >= 0; i-- {
		popNode := stack[i]
		ans.Next = popNode
	}
	return stack[len(stack)-1]
}

func reverseListV3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func display(data *ListNode) []int {
	var ans []int

	for data != nil {
		ans = append(ans, data.Val)
		data = data.Next
	}
	return ans
}
func main() {
	data := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Println(display(data))
	reverseListV1(data)

	fmt.Println(display(data))
}
