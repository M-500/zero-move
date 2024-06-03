package _7_fast_sort

import "testing"

// @Description
// @Author 代码小学生王木木

func quickSort(arr []int, L, R int) {
	if L < R {
		temp := partition(arr, L, R)
		quickSort(arr, L, temp[0]-1)
		quickSort(arr, temp[1], R)
	}
}
func partition(arr []int, L, R int) []int {
	less := L - 1
	more := R + 1
	for L < more {
		if arr[L] < arr[R] {
			arr[L], arr[less+1] = arr[less+1], arr[L]
			less++
			L++
		} else if arr[L] > arr[R] {
			arr[L], arr[more-1] = arr[more-1], arr[L]
			more--
		} else {
			L++
		}
	}
	// 返回等于target的最小下标和最大下标
	return []int{less + 1, more - 1}
}

func TestQuickSort(t *testing.T) {
	arr1 := []int{3, 7, 5, 2, 1, 4, 8, 12, 1, 2, 3, 122}
	quickSort(arr1, 0, len(arr1)-1)
	t.Log(arr1)
}
