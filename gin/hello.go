package main

import (
	"github.com/gin-gonic/gin"
)

//Restaurant ....test
type Restaurant struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	result := Restaurant{
		ID:   3,
		Name: "XXXXX",
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, result)
	})
	r.GET("/gin", func(c *gin.Context) {
		c.JSON(200, Restaurant{ID: 1, Name: "gin json"})
	})
	r.Run(":8080")
}
