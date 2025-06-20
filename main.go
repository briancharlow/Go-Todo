package main

import (
	"github.com/gin-gonic/gin"
	"todo.com/go/api"
)

func main() {
	router := gin.Default()

	// Routes
	router.POST("/create", api.CreateTodo)
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	router.GET("/todos", api.GetTodos)

	// Start the server
	err := router.Run(":8080") // Starts the server on port 8080
	if err != nil {
		panic(err)
	}
}
