package base

import "sync"

// @Description
// @Author 代码小学生王木木

type node struct {
	Name   string
	Weight int // 代表权重
	Host   string
	Port   int
}

type WeightRoundRobinBalancer struct {
	nodeList []*node
	lock     sync.Mutex
}

func (w *WeightRoundRobinBalancer) weightRoundRobin() *Node {
	w.lock.Lock()
	defer w.lock.Unlock()

	return nil
}
