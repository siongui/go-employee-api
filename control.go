package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// getEmployees responds with the list of all employees as JSON.
func getEmployees(c *gin.Context) {
	employees, err := SelectAllEmployees()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Error(err)
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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		log.Error(err)
		return
	}

	// TODO: check if id already exists in database

	// Add the new employee to the database.
	if _, err := InsertEmployee(newEmployee); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Error(err)
		return
	}
	c.IndentedJSON(http.StatusCreated, newEmployee)
	log.Info(newEmployee, " inserted to database")
}

// getEmployeeByID locates the employee whose ID value matches the id
// parameter sent by the client, then returns that employee as a response.
func getEmployeeByID(c *gin.Context) {
	s := c.Param("id")

	id, err := strconv.Atoi(s)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "employee not found"})
		log.Error(err)
		return
	}

	// Looking for an employee whose ID value matches the parameter in the
	// database.
	e, err := SelectById(id)

	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "employee not found"})
		log.Error(err)
		return
	}

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, e)
	log.Info(e, " selected and returned")
}

// deleteEmployeeByID locates the employee whose ID value matches the id
// parameter sent by the client.
func deleteEmployeeByID(c *gin.Context) {
	s := c.Param("id")

	id, err := strconv.Atoi(s)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "employee not found"})
		log.Error(err)
		return
	}

	_, err = DeleteById(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "employee with id " + s + " deleted"})
	log.Info("employee with id " + s + " deleted")
}

// updateEmployee updates an employee from JSON received in the request body.
func updateEmployee(c *gin.Context) {
	var e Employee

	// Call BindJSON to bind the received JSON to e.
	if err := c.BindJSON(&e); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		log.Error(err)
		return
	}

	// Update the employee to the database.
	if _, err := UpdateEmployee(e); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Error(err)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "employee with id " + strconv.Itoa(e.Id) + " updated"})
	log.Info(e, " updated in database")
}
