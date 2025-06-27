package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	resp, err := http.Get("http://localhost:3001/todos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}
	defer resp.Body.Close()

	var todos []Todo
	if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid JSON response"})
		return
	}

	c.JSON(http.StatusOK, todos)
}
