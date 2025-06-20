package api

import (
	"net/http"
	"todo.com/go/model"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	todos := model.GetTodos()
	c.JSON(http.StatusOK, todos)
}
