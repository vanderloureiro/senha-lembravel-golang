package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vanderloureiro/senha-lembravel-golang/data"
)

func main() {

	fmt.Println("Hello World!")
	r := gin.Default()
	r.GET("/", retornaMensagem)
	r.Run()
}

func retornaMensagem(c *gin.Context) {
	value := data.GetRandomWord()
	c.JSON(200, gin.H{
		"message": value,
	})
}
