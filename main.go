package main

import (
	"net/http"
	"strconv"

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

// getEmployeeByID locates the employee whose ID value matches the id
// parameter sent by the client, then returns that employee as a response.
func getEmployeeByID(c *gin.Context) {
	s := c.Param("id")

	id, err := strconv.Atoi(s)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "employee not found"})
		return
	}

	// Loop over the list of employees, looking for
	// an employee whose ID value matches the parameter.
	for _, e := range employees {
		if e.Id == id {
			c.IndentedJSON(http.StatusOK, e)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "employee not found"})
}

func main() {
	router := gin.Default()
	router.GET("/employees", getEmployees)
	router.GET("/employee/:id", getEmployeeByID)
	router.POST("/employee", postEmployee)

	router.Run("localhost:8080")
}
