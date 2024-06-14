package base

import "sync/atomic"

// @Description
// @Author 代码小学生王木木

// 手写轮询负载均衡

type Node struct {
	Name string
	Host string
	Port int
}

type RoundBalancer struct {
	nodeList []*Node
	idx      *atomic.Int32
}

func (r *RoundBalancer) roundRobin() *Node {
	idx := int(r.idx.Add(1))
	i := len(r.nodeList)
	return r.nodeList[idx%i]
}
