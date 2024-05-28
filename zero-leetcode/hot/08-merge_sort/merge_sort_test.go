package _8_merge_sort

// @Description
// @Author 代码小学生王木木

// 递归实现归并算法
func sortArray(nums []int) []int {
	merge := func(left, right []int) []int {
		res := make([]int, len(left)+len(right))
		var l, r, i int
		// 通过遍历完成比较填入res中
		for l < len(left) && r < len(right) {
			if left[l] <= right[r] {
				res[i] = left[l]
				l++
			} else {
				res[i] = right[r]
				r++
			}
			i++
		}
		// 如果left或者right还有剩余元素，添加到结果集的尾部
		copy(res[i:], left[l:])
		copy(res[i+len(left)-l:], right[r:])
		return res
	}
	var sort func(nums []int) []int
	sort = func(nums []int) []int {
		if len(nums) <= 1 {
			return nums
		}
		// 拆分递归与合并
		// 分割点
		mid := len(nums) / 2
		left := sort(nums[:mid])
		right := sort(nums[mid:])
		return merge(left, right)
	}
	return sort(nums)
}
