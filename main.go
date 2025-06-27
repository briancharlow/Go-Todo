package main

import (
	"github.com/gin-gonic/gin"
	"todo.com/go/api"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	router.GET("/todos", api.GetTodos)
	router.POST("/todos", api.CreateTodo)
	router.PUT("/todos/:id", api.UpdateTodo)
	router.DELETE("/todos/:id", api.DeleteTodo)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
