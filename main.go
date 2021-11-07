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

// getEmployees responds with the list of all employees as JSON.
func getEmployees(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, employees)
}

// postEmployee adds an employee from JSON received in the request body.
func postEmployee(c *gin.Context) {
	var newEmployee Employee

	// Call BindJSON to bind the received JSON to
	// newEmployee.
	if err := c.BindJSON(&newEmployee); err != nil {
		return
	}

	// Add the new employee to the slice.
	employees = append(employees, newEmployee)
	c.IndentedJSON(http.StatusCreated, newEmployee)
}

func main() {
	router := gin.Default()
	router.GET("/employees", getEmployees)
	router.POST("/employee", postEmployee)

	router.Run("localhost:8080")
}
