//@Author: wulinlin
//@Description:
//@File:  redis_lock_test
//@Version: 1.0.0
//@Date: 2024/05/26 13:20

package redis_lock

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"zero-template/cache/redis-lock/mocks"
)

// mockgen -package=mocks -destination=cache/redis-lock/mocks/redis_cmdable.mock.go github.com/redis/go-redis/v9 Cmdable
// https://blog.csdn.net/weixin_57135751/article/details/128825424
func TestClientLock(t *testing.T) {
	testCase := []struct {
		name string
		mock func(ctrl *gomock.Controller) redis.Cmdable
		key  string

		wantErr error
		wantRes *Lock
	}{
		{
			name: "setnx 命令执行错误",
			key:  "key1",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := mocks.NewMockCmdable(ctrl)
				res := redis.NewBoolResult(false, context.DeadlineExceeded)
				cmd.EXPECT().SetNX(context.Background(), "key1", gomock.Any(), time.Minute).Return(res)
				return cmd
			},
			wantErr: context.DeadlineExceeded,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			client := NewClient(tc.mock(controller))
			lock, err := client.TryLock(context.Background(), tc.key, time.Minute)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantRes.key, lock)
			assert.NotEmpty(t, lock.value)
			// notEmpty
		})
	}

}
