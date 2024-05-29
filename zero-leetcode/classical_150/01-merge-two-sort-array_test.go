package classical_150

import (
	"fmt"
	"testing"
)

// @Description
// @Author 代码小学生王木木

func Test01(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3

	fmt.Println(nums1)
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 && n == 0 {
		return
	}
	for i := m + n - 1; i > 0; i-- {
		if nums1[m-1] > nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
			continue
		}
		nums1[i] = nums2[n-1]
		n--
	}
}

// mergeAC
//
//	@Description: 考虑了特殊情况的
//	@param nums1
//	@param m
//	@param nums2
//	@param n
func mergeAC(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	tail := m + n - 1
	for p1 >= 0 || p2 >= 0 {
		var cur int
		// 处理临界值的情况
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
		tail--
	}
}
