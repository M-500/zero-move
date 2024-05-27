package limit

import "sync/atomic"

// @Description
// @Author 代码小学生王木木

type CounterLimiter struct {
	cnt       *atomic.Int32
	threshold int32
}

func (c *CounterLimiter) Limit() {
	cnt := c.cnt.Add(1)
	defer func() {
		c.cnt.Add(-1)
	}()
	if cnt > c.threshold {

		return
	}
	// 调用回调函数
}
