package _1_int_to_str

import (
	"fmt"
	"strings"
	"testing"
)

// @Description
// @Author 代码小学生王木木

func intToStr(a uint) string {
	if a == 0 {
		return "零"
	}
	numHash := map[int]string{
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
	steps := []string{"", "十", "百", "千", "万", "十万", "百万", "千万", "亿"}
	oldLen := len(fmt.Sprintf("%d", a)) // 数字有多少位
	res := ""
	numArr := make([]int, oldLen)    // 阿拉伯数组
	strArr := make([]string, oldLen) // 中文的数组
	i := 0
	for a > 0 {
		k := a % 10 // 最后一位
		numArr[oldLen-1-i] = int(k)
		strArr[oldLen-1-i] = numHash[int(k)]
		i++
		a = a / 10
	}
	// numArr [1 0 0 0 8]
	// strArr [一 零 零 零 八]
	zeroFlag := false
	for idx, item := range numArr {
		if item == 0 {
			// 第一次遇到了零
			zeroFlag = true
		} else {
			// 连续的零 只会添加一个零
			if zeroFlag {
				res += "零"
			}
			res += strArr[idx]          // 添加中文数字
			step := steps[oldLen-idx-1] // 获取当前的单位
			res += step                 // 添加单位
			zeroFlag = false
		}

	}
	return res
}

// 定义中文数字和单位
var numMap = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
var unitMap = []string{"", "十", "百", "千", "万", "十", "百", "千", "亿"}

func numberToChinese(num int) string {
	if num == 0 {
		return "零"
	}
	// 将数字转换为字符串
	numStr := fmt.Sprintf("%d", num)
	numLen := len(numStr)

	// 用于存储结果
	var result strings.Builder

	zeroFlag := false
	for i, digit := range numStr {
		// 当前数字
		n := digit - '0'
		// 当前单位
		unit := unitMap[numLen-1-i]

		if n == 0 {
			// 连续零的处理
			zeroFlag = true
		} else {
			if zeroFlag {
				// 添加一个零
				result.WriteString(numMap[0])
				zeroFlag = false
			}
			// 添加数字和单位
			result.WriteString(numMap[n])
			result.WriteString(unit)
		}

		// 处理万和亿的单位
		if (numLen-1-i)%4 == 0 && (numLen-1-i) != 0 {
			result.WriteString(unitMap[(numLen-1-i)/4+3])
		}
	}

	// 去掉末尾的"零"
	return result.String()
}
func TestDemo(t *testing.T) {
	t.Log(intToStr(120))
	t.Log(intToStr(12018))
	//t.Log(intToStr(1999))
	//t.Log(intToStr(123001))
	//t.Log(numberToChinese(108))
	//t.Log(numberToChinese(10008))
	//t.Log(numberToChinese(10008))
	//t.Log(numberToChinese(199999999))
}
