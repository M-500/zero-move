package _7_fast_sort

import "testing"

// @Description
// @Author 代码小学生王木木

func TestBaseV2(t *testing.T) {
	arr1 := []int{3, 7, 5, 2, 1, 4, 8, 12, 1, 2, 3, 122}
	//Color(arr1, 4)
	Partition(arr1, 0, len(arr1)-1, 4)
	t.Log(arr1)
}

func Color(arr []int, target int) {
	arrLen := len(arr)
	if arrLen == 0 {
		return
	}
	leftIndex := -1      // 记录小于目标target的
	rightIndex := arrLen // 记录大于目标target的下标
	for i := 0; i < arrLen; {
		// 如果当前遍历因子和右边界重合，那么循环结束
		if rightIndex == i {
			break
		}
		if arr[i] == target {
			i++
			continue
		}
		if arr[i] < target {
			arr[i], arr[leftIndex+1] = arr[leftIndex+1], arr[i]
			leftIndex++
			i++
			continue
		}
		if arr[i] > target {
			arr[i], arr[rightIndex-1] = arr[rightIndex-1], arr[i]
			rightIndex--
		}
	}
}

func Partition(arr []int, L, R, target int) []int {
	less := L - 1
	more := R + 1
	for L < more {
		if arr[L] < target {
			arr[L], arr[less+1] = arr[less+1], arr[L]
			less++
			L++
		} else if arr[L] > target {
			arr[L], arr[more-1] = arr[more-1], arr[L]
			more--
		} else {
			L++
		}
	}
	// 返回等于target的最小下标和最大下标
	return []int{less + 1, more - 1}
}
