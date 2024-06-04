package _9_base_sort

import "testing"

// @Description
// @Author 代码小学生王木木

func BubbleSort(arr []int) {
	var arrLen = len(arr)
	if arrLen <= 1 {
		return
	}
	for i := 0; i < arrLen; i++ {
		for j := i + 1; j < arrLen; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	arr := []int{3, 7, 1, 4, 5, 9, 2, 8}
	BubbleSort(arr)
	t.Log(arr)

}
