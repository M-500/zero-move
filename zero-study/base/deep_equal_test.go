//@Author: wulinlin
//@Description:
//@File:  deep_equal_test
//@Version: 1.0.0
//@Date: 2024/06/09 10:53

package base

import (
	"reflect"
	"testing"
)

func TestArrayComparable(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	t.Log(a == b)
}

func TestSliceComparable(t *testing.T) {
	a := [3]int{1, 2, 3}
	//b := []int{1, 2, 3}
	ref := reflect.TypeOf(a)
	ref.Comparable()
	t.Log(ref.Comparable())
	//t.Log(a == b)
}

func TestIsComparable(t *testing.T) {
	var a int64 = 1
	of := reflect.TypeOf(a)
	t.Log(of.Comparable())
}
