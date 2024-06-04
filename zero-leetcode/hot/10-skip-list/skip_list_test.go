package _0_skip_list

import (
	"math/rand"
	"testing"
)

const (
	maxLevel = 32   // 最大层高32层
	pFactor  = 0.25 // 随机阈值  根据这个阈值决定是否要上层开开辟索引
)

// Node
// @Description: 节点
type Node struct {
	val     int
	forward []*Node
}

// SkipList
// @Description: 跳表
type SkipList struct {
	head  *Node // 头节点
	level int   // 层高
}

func Constructor() SkipList {
	return SkipList{
		head: &Node{
			val:     -1,
			forward: make([]*Node, maxLevel),
		},
		level: 0,
	}
}

func (s SkipList) randomLevel() int {
	lv := 1
	for lv < maxLevel && rand.Float64() < pFactor {
		lv++
	}
	return lv
}

func (this *SkipList) Search(target int) bool {
	cur := this.head
	// 从最高层往下，逐层寻找
	for i := this.level - 1; i >= 0; i-- {
		// 找到I层小于且最接近 目标值的元素
		for cur.forward[i] != nil && cur.forward[i].val < target {
			cur = cur.forward[i]
		}
	}
	cur = cur.forward[0]
	// 检测当前元素的值是否等于 target
	return cur != nil && cur.val == target
}

func (this *SkipList) Add(num int) {
	nodes := make([]*Node, maxLevel)
	for i := range nodes {
		nodes[i] = this.head
	}
	curr := this.head
	for i := this.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 num 的元素
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		nodes[i] = curr
	}
	lv := this.randomLevel()
	this.level = max(this.level, lv)
	// 新增一个节点
	newNode := &Node{num, make([]*Node, lv)}
	for i, node := range nodes[:lv] {
		// 对第 i 层的状态进行更新，将当前元素的 forward 指向新的节点
		newNode.forward[i] = node.forward[i]
		node.forward[i] = newNode
	}
}

func (this *SkipList) Erase(num int) bool {
	update := make([]*Node, maxLevel)
	curr := this.head
	for i := this.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 num 的元素
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	curr = curr.forward[0]
	// 如果值不存在则返回 false
	if curr == nil || curr.val != num {
		return false
	}
	for i := 0; i < this.level && update[i].forward[i] == curr; i++ {
		// 对第 i 层的状态进行更新，将 forward 指向被删除节点的下一跳
		update[i].forward[i] = curr.forward[i]
	}
	// 更新当前的 level
	for this.level > 1 && this.head.forward[this.level-1] == nil {
		this.level--
	}
	return true
}

/**
 * Your SkipList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */

func TestSkipList(t *testing.T) {
	obj := Constructor()

	obj.Add(1)
	obj.Add(5)
	obj.Add(7)
	obj.Add(9)
	param_1 := obj.Search(3)
	t.Log(param_1)
	t.Log(obj.Search(5))
	obj.Add(6)
	obj.Erase(5)
}
