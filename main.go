package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Buy milk", Completed: false},
	{ID: "2", Item: "Buy eggs", Completed: false},
	{ID: "3", Item: "Buy bread", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func createTodo(context *gin.Context) {
	var todo todo
	context.BindJSON(&todo)
	todos = append(todos, todo)
	context.JSON(http.StatusCreated, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:8080")
}
