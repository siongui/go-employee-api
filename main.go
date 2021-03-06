package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	InitSQLite(true)
	_, err := CreateEmployeeTable()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/employees", getEmployees)
	router.GET("/employee/:id", getEmployeeByID)
	router.POST("/employee", postEmployee)
	router.DELETE("/employee/:id", deleteEmployeeByID)
	router.PUT("/employee", updateEmployee)

	router.Run("localhost:8080")
}
