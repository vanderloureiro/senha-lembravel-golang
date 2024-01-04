package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Hello World!")
	r := gin.Default()
	r.GET("/", retornaMensagem)
	r.Run()
}

func retornaMensagem(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Senha Lembrável API",
	})
}
