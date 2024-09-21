package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func main() {
	// router is the server
	router := gin.Default()

	// now create end points
	router.GET("/todos", getTodos)

	router.POST("add", addTodo)
	router.POST("update/:id", updateTodoByID)
	router.PATCH("patch/:id", toggleStatus)
	router.GET("/todos/:id", getTodo)

	router.Run("localhost:9090")

}

func getTodos(context *gin.Context) { // this content going to contain bunch of information about incoming http request
	// return todos into json

	context.IndentedJSON(http.StatusOK, todos)

}

func addTodo(context *gin.Context) {
	var newTodo todo
	err := context.BindJSON(&newTodo)
	if err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, todos)

}

func getTodo(context *gin.Context) {

	id := context.Param("id")
	todo, err := getTodoByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	context.IndentedJSON(http.StatusCreated, todo)
}

// returns either todo or new error
func getTodoByID(id string) (*todo, error) {

	for i, t := range todos {

		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("Todo not found")
}

func updateTodoByID(context *gin.Context) {

	id := context.Param("id")
	var updatedTodo todo
	existingTodo, error := getTodoByID(id)
	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No id found to update"})
		return
	}
	print(existingTodo)

	err := context.BindJSON(&updatedTodo)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	for i, t := range todos {

		if t.ID == id {
			todos[i] = updatedTodo
		}
	}

	context.IndentedJSON(http.StatusOK, updatedTodo)
}

func toggleStatus(context *gin.Context) {

	id := context.Param("id")

	existingTodo, err := getTodoByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No id found to toggle"})
		return
	}

	existingTodo.Completed = !existingTodo.Completed

	context.IndentedJSON(http.StatusOK, existingTodo)
}
