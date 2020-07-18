package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// "net/http"
)

func main() {
	fmt.Println("hello")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()

}
