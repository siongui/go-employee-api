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

	// test insertion
	_, err = InsertEmployee(employees[0])
	if err != nil {
		t.Error(err)
		return
	}

	// test insertion
	_, err = InsertEmployee(employees[1])
	if err != nil {
		t.Error(err)
		return
	}

	// test reading all employees
	es, err := SelectAllEmployees()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(es)

	// test selection by id
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

	// test deletion by id
	_, err = DeleteById(1)
	if err != nil {
		t.Error(err)
		return
	}
	es, err = SelectAllEmployees()
	if err != nil {
		t.Error(err)
		return
	}
	if len(es) != 1 {
		t.Error("should left only one record after deletion")
		return
	}
	t.Log(es)
}
