package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vanderloureiro/senha-lembravel-golang/pkg/generator"
)

func main() {

	fmt.Println("Hello World!")
	r := gin.Default()
	r.GET("/", retornaMensagem)
	r.Run()
}

func retornaMensagem(c *gin.Context) {
	value := generator.GeneratePassword()
	c.JSON(200, gin.H{
		"message": value,
	})
}
