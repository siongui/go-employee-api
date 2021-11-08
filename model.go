package main

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

// Employee represents the structure of a employee.
type Employee struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

var db *bun.DB
var ctx = context.Background()

// CreateEmployeeTable creates table in the database to store employee data.
func CreateEmployeeTable() (sql.Result, error) {
	return db.NewCreateTable().Model((*Employee)(nil)).Exec(ctx)
}

// InitSQLite initialize in-memory database to store data. The verbose flag
// indicates whether to print all queries to stdout.
func InitSQLite(verbose bool) (err error) {
	// Open an in-memory SQLite database.
	sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		return
	}

	sqldb.SetMaxOpenConns(1)

	// Create a Bun db on top of it.
	db = bun.NewDB(sqldb, sqlitedialect.New())

	// If you are using an in-memory database, you need to configure *sql.DB
	// to NOT close active connections. Otherwise, the database is deleted
	// when the connection is closed.
	//sqldb.SetMaxIdleConns(1000)
	//sqldb.SetConnMaxLifetime(0)

	// Print all queries to stdout.
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(verbose)))

	return
}

// SelectAllEmployees reads all employees from the database.
func SelectAllEmployees() (es []Employee, err error) {
	err = db.NewSelect().
		Model(&es).
		OrderExpr("id ASC").
		Scan(ctx)
	return
}

// InsertEmployee inserts one employee in the database.
func InsertEmployee(e Employee) (result sql.Result, err error) {
	return db.NewInsert().Model(&e).Exec(ctx)
}

// SelectById selects the record by id in the database.
func SelectById(id int) (e Employee, err error) {
	err = db.NewSelect().
		Model(&e).
		Where("id = ?", id).
		Limit(1).
		Scan(ctx)
	return
}
