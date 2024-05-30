package gin_stu

// @Description
// @Author 代码小学生王木木

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestGin01(t *testing.T) {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "哈哈哈哈",
		})
		return
	})

	router.Run(":8090")
}
