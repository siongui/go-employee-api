package main

import (
	"database/sql"
	"testing"
)

var employees = []Employee{
	{Id: 1, Name: "Hello", Title: "Engineer"},
	{Id: 2, Name: "World", Title: "Manager"},
}

func TestSqliteOperation(t *testing.T) {
	InitSQLite(true)

	_, err := CreateEmployeeTable()
	if err != nil {
		t.Error(err)
		return
	}

	_, err = InsertEmployee(employees[0])
	if err != nil {
		t.Error(err)
		return
	}

	_, err = InsertEmployee(employees[1])
	if err != nil {
		t.Error(err)
		return
	}

	es, err := SelectAllEmployees()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(es)

	e, err := SelectById(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(e)

	_, err = SelectById(777)
	if err != sql.ErrNoRows {
		t.Error(err)
		return
	}
}
