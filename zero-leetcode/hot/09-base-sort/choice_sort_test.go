package _9_base_sort

import "testing"

// @Description
// @Author 代码小学生王木木

func ChoiceSort(arr []int) {
	arrLen := len(arr)
	if arrLen < 2 {
		return
	}
	for i := 0; i < arrLen; i++ {
		temp := i
		for j := i; j < arrLen; j++ {
			if arr[temp] > arr[j] {
				temp = j
			}
		}
		if temp != i {
			arr[i], arr[temp] = arr[temp], arr[i]
		}
	}
}

func TestChoiceSort(t *testing.T) {
	arr := []int{3, 7, 1, 4, 5, 9, 2, 8}
	ChoiceSort(arr)
	t.Log(arr)

}
