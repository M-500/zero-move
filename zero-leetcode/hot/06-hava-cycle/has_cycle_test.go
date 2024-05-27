package _6_hava_cycle

// @Description
// @Author 代码小学生王木木

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
func hasCycleHash(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	hashMap := make(map[*ListNode]bool, 0)
	for head != nil {
		_, ok := hashMap[head]
		if ok {
			return true
		}
		hashMap[head] = true
		head = head.Next
	}
	return false
}
