package hh

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	msgApiV1 := r.Group("/api/v1/msg")

	msgApiV1.GET("/ping", func(c *gin.Context) {
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
