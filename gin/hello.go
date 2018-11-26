package main

import (
	"github.com/gin-gonic/gin"
)

type Restaurant struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	result := Restaurant{
		Id:   3,
		Name: "XXXXX",
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, result)
	})
	r.GET("/gin", func(c *gin.Context) {
		c.JSON(200, Restaurant{Id: 1, Name: "gin json"})
	})
	r.Run(":8080")
}
