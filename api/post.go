package api

import (
	"net/http"
	"todo.com/go/model"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todo model.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	model.AddTodo(todo)

	response := model.Response{
		Message: "Todo created successfully",
		Todo:    todo,
	}

	c.JSON(http.StatusCreated, response)
}
