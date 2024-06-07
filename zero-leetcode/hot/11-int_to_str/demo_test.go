package _1_int_to_str

import (
	"strconv"
	"testing"
)

// @Description
// @Author 代码小学生王木木

var chMap = map[int]string{
	0: "零",
	1: "一",
	2: "二",
	3: "三",
	4: "四",
	5: "五",
	6: "六",
	7: "七",
	8: "八",
	9: "九",
}

var ss = []string{"", "十", "百", "千", "万", "十", "百", "千", "亿"}

func demo(a int) string {
	if a == 0 {
		return "零"
	}
	var ans = ""
	// 1. 切分出来一个个数字
	arrLen := len(strconv.Itoa(a))
	numSlice := make([]int, arrLen)   // 1,2,3 栈
	chSlice := make([]string, arrLen) //一 二 三

	// 2. 映射
	for i := arrLen - 1; a > 0; i-- {
		splitData := a % 10 // 3  2 1
		numSlice[i] = splitData
		chSlice[i] = chMap[splitData]
		a = int(a / 10) // 12  1 0
	}

	// 3. 拼接
	zeroFlg := false
	for i, j := range numSlice {
		if j == 0 {
			zeroFlg = true
		} else {
			if zeroFlg {
				ans += "零"
			}
			ans += chSlice[i] // 拼接数字
			unit := ss[arrLen-i-1]
			ans += unit // 拼接单位
			zeroFlg = false
		}

	}
	return ans
}

func TestHahahah(t *testing.T) {
	t.Log(demo(100423120))
}
