package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

var employees = []Employee{
	{Id: 1, Name: "Hello", Title: "Engineer"},
	{Id: 2, Name: "World", Title: "Manager"},
}

func getEmployees(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, employees)
}

func main() {
	router := gin.Default()
	router.GET("/employees", getEmployees)

	router.Run("localhost:8080")
}
