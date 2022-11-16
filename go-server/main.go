package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID   string
	Todo string
	Done bool
}

var todoList = []todo{
	{ID: "1", Todo: "Take a walk.", Done: false},
	{ID: "2", Todo: "Drink water!", Done: false},
	{ID: "3", Todo: "Have breakfast.", Done: false},
}

func welcome(ctx *gin.Context) {
	welcomeMessage := "Congratulations!"
	ctx.IndentedJSON(http.StatusOK, welcomeMessage)
}

func getAllTodos(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, todoList)
}

func getTodoById(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, todo := range todoList {
		if todo.ID == id {
			ctx.IndentedJSON(http.StatusOK, todo)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func addTodo(ctx *gin.Context) {
	// parse the POST body
	var newTodo todo

	if err := ctx.BindJSON(&newTodo); err != nil {
		return
	}

	todoList = append(todoList, newTodo)
	ctx.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	// initializes gin router
	router := gin.Default()
	router.GET("/", welcome)
	router.GET("/todos", getAllTodos)
	router.GET("/todo/:id", getTodoById)
	router.POST("/todo", addTodo)
	router.Run("localhost:8080")
}
