package base

import (
	"strconv"
	"strings"
	"testing"
)

// @Description
// @Author 代码小学生王木木
const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

func TestDemo1(t *testing.T) {

	//t.Log("草了mutexLocked", mutexLocked)
	//t.Log("草了mutexWoken", mutexWoken)
	//t.Log("草了mutexStarving", mutexStarving)
	//t.Log("草了mutexWaiterShift", mutexWaiterShift)
	for i, i2 := range "123456" {
		t.Log(i, i2)
	}
	t.Log(intToString(1123486))
	t.Log(intToString(10806))

}

var chMap = map[uint]string{
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
var DW = []string{"", "十", "百", "千", "万", "十", "百", "千", "亿"}

func intToString(a uint) string {
	if a == 0 {
		return "零"
	}
	// 1. 获取a的位数
	aLen := len(strconv.Itoa(int(a)))
	numArr := make([]int, aLen)      // 类似于 [1, 0, 0, 0 ,8, 6]
	chinaArr := make([]string, aLen) // 类似于 ["一", "零", 。。。]
	var ans strings.Builder
	for i := aLen - 1; a > 0; i-- {
		temp := a % 10
		numArr[i] = int(temp)
		chinaArr[i] = chMap[temp]
		a = a / 10
	}
	// 开始拼接字符串
	zeroFlg := false
	for i, j := range numArr {
		if j == 0 {
			zeroFlg = true
		} else {
			if zeroFlg {
				ans.WriteString("零")
			}
			// 1. 获取单位
			tp := aLen - 1 - i
			curDW := DW[tp]              // 获取单位
			ans.WriteString(chinaArr[i]) // 写入数字
			ans.WriteString(curDW)       // 写入单位
			if tp%4 == 0 && tp != 0 {
				//ans.WriteString(DW[tp/4+3]) // 写入单位
			}
			tp--
			zeroFlg = false // 肯定不是连续的0了
		}

	}
	return ans.String()
}
