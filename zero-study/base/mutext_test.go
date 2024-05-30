//@Author: wulinlin
//@Description:
//@File:  mutext_test
//@Version: 1.0.0
//@Date: 2024/05/29 20:45

package base

import (
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
	lock := sync.Mutex{}
	lock.Lock()
}
