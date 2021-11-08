package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// getEmployees responds with the list of all employees as JSON.
func getEmployees(c *gin.Context) {
	employees, err := SelectAllEmployees()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database operation error"})
		return
	}

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

	// Add the new employee to the database.
	if _, err := InsertEmployee(newEmployee); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database operation error"})
		return
	}
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

	// Looking for an employee whose ID value matches the parameter in the
	// database.
	e, err := SelectById(id)

	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "employee not found"})
		return
	}

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database operation error"})
		return
	}

	c.IndentedJSON(http.StatusOK, e)
}
