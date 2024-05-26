//@Author: wulinlin
//@Description:
//@File:  redis_lock
//@Version: 1.0.0
//@Date: 2024/05/24 22:13

package redis_lock

import (
	"context"
	_ "embed"
	"errors"
	"github.com/nu7hatch/gouuid"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	errFailedToPreemptLock = errors.New("redis-lock: 抢锁失败")
	errLockNotExist        = errors.New("redis-lock:锁不存在")
	errLockNotHold         = errors.New("redis-lock:没有持有锁")
	//go:embed script/unlock.lua
	unlockLua string

	//go:embed script/refresh.lua
	refreshLua string
)

type Client struct {
	client redis.Cmdable
}

func NewClient(client redis.Cmdable) *Client {
	return &Client{client: client}
}

type Lock struct {
	key      string
	value    string
	client   redis.Cmdable
	expire   time.Duration
	stopChan chan struct{}
}

// AutoRefresh
//
//	@Description:
//	@receiver l
//	@param interval  续约的间隔是时间
//	@param timeout  续约一次的超时时间
func (l *Lock) AutoRefresh(interval time.Duration, timeout time.Duration) error {
	timeoutChan := make(chan struct{}, 1)
	// 间隔多久续约一次
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			err := l.Refresh(ctx, l.expire)
			cancel()
			if err == context.DeadlineExceeded {
				timeoutChan <- struct{}{}
				continue
			}
			if err != nil {
				return err
			}
		case <-timeoutChan:
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			err := l.Refresh(ctx, l.expire)
			cancel()
			if err == context.DeadlineExceeded {
				timeoutChan <- struct{}{}
				continue
			}
			if err != nil {
				return err
			}
		case <-l.stopChan:
			return nil
		}
	}
	return nil
}
func (l *Lock) Refresh(ctx context.Context, expire time.Duration) error {
	result, err := l.client.Eval(ctx, refreshLua, []string{l.key}, l.value, expire).Int64()
	// 如果命令执行出错，可能会进入下面的分支
	if errors.Is(err, redis.Nil) {
		return errLockNotHold
	}
	if err != nil {
		return err
	}
	if result != 1 {
		return errLockNotExist
	}
	return nil
}

// Unlock
//
//	@Description: 这个才优雅，导致了只有持有锁的人才有资格释放锁
//	@receiver l
func (l *Lock) Unlock(ctx context.Context) error {
	// 1. 判断这把锁是不是我加的锁
	//res, err := l.client.Get(ctx, l.key).Result()
	//if err != nil {
	//	return err
	//}
	//if res != l.value {
	//	// 说明锁不是自己的锁  这里不能这么写！！
	//	return err
	//}// 这么写有问题，并发不安全

	//result, err := l.client.Del(ctx, l.key).Result()
	//if err != nil {
	//	return err
	//}
	//if result != 1 {
	//	// 代表你加的锁 过期了
	//	return errLockNotExist
	//} // 这么写也有问题，很容易释放掉不属于自己的锁

	defer func() {
		close(l.stopChan)
	}()
	result, err := l.client.Eval(ctx, unlockLua, []string{l.key}, l.value).Int64()
	// 如果命令执行出错，可能会进入下面的分支
	if errors.Is(err, redis.Nil) {
		return errLockNotHold
	}
	if err != nil {
		return err
	}
	if result != 1 {
		return errLockNotExist
	}
	return nil
}

func (c *Client) TryLock(ctx context.Context, key string, expiration time.Duration) (*Lock, error) {

	uuidBt, err := uuid.NewV4() // 用来唯一标识系铃人
	if err != nil {
		return nil, err
	}
	ok, err := c.client.SetNX(ctx, key, uuidBt.String(), expiration).Result()
	if err != nil {
		return nil, err
	}
	if !ok {
		// 别人获取了锁，你抢锁失败
		return nil, errFailedToPreemptLock
	}
	return &Lock{
		client: c.client,
		key:    key,
		value:  uuidBt.String(),
		expire: expiration,
	}, err
}

// TODO 这个写法不优雅
func (c *Client) Unlock(ctx context.Context, lock *Lock) {

}
