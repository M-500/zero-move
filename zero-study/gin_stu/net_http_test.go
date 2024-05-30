package gin_stu

import (
	"net/http"
	"testing"
)

// @Description
// @Author 代码小学生王木木

func TestNetHTTP01(t *testing.T) {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	// 开启一个HTTP服务，在8091端口上
	http.ListenAndServe(":8091", nil)
}
