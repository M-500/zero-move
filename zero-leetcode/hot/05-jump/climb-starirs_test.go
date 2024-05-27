package _5_jump

// @Description
// @Author 代码小学生王木木

func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairsV1(n int) int {
	if n < 2 {
		return n
	}
	var climbArray = make([]int, n)
	climbArray[0] = 1
	climbArray[1] = 2
	for i := 2; i < n; i++ {
		climbArray[i] = climbArray[i-1] + climbArray[i-2]
	}
	return climbArray[n-1]
}

func climbStairsV2(n int) int {
	if n <= 2 {
		return n
	}
	slow, fast := 1, 2
	var ans int
	for i := 2; i < n; i++ {
		ans = slow + fast
		slow = fast
		fast = ans
	}
	return ans
}
