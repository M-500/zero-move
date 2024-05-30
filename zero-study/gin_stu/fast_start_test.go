//@Author: wulinlin
//@Description:
//@File:  fast_start_test
//@Version: 1.0.0
//@Date: 2024/05/29 20:53

package gin_stu

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestGin01(t *testing.T) {
	route := gin.Default()
	route.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hahah",
		})
	})
	route.Run(":8090")
}
