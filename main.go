package main

import (
	"github.com/dinudk02/golang/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVar()

}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong1234",
		})
	})
	r.Run()
}
