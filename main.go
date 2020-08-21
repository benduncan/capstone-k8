package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

/*
Project Capstone
Golang / Jenkins / AWS EKS example
*/

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": SayHello(),
		})
	})

	r.GET("/capstone", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"payload": "Special payload for project Capstone ...",
		})
	})

	r.Run(":8000")
}

func SayHello() (welcome string) {

	return fmt.Sprintf("%s", "Yo world!")
}
