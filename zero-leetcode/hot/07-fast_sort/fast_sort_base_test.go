package _7_fast_sort

import "testing"

// @Description 荷兰三色国旗问题 给定一个数组arr，和一个数num，请把小于等于num的数放在数组的左边，大于num的数放在数组的右边。 要求额外空间复杂度O(1)，时间复杂度O(N) 不要求左半部分或者右半部分有序，只要保证target小于左边就好了
// @Author 代码小学生王木木

func TestThreeColor(t *testing.T) {
	arr := []int{3, 7, 5, 2, 1, 4, 8}
	ThreeColorV1(arr, 4)
	t.Log(arr)
	arr1 := []int{3, 7, 5, 2, 1, 4, 8, 12, 1, 2, 3, 122}
	ThreeColorV1(arr1, 4)
	t.Log(arr1)
}

func ThreeColorV1(arr []int, target int) []int {
	arrLen := len(arr)
	if arrLen <= 1 {
		return arr
	}
	x := -1 // 用于记录第一个大于等于目标的下表
	for i := 0; i < arrLen; i++ {
		if arr[i] <= target {
			arr[i], arr[x+1] = arr[x+1], arr[i]
			x++
			continue
		}
	}
	return arr
}
