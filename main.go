package main

import (
	"PluckyAPI/Controllers"
	"PluckyAPI/Utilities/dbhelpers"
	"database/sql"
	"fmt"
)

// TODO: 1. Sort out opening a connection to the database
//       2. Decide on format of incoming json structure
//       3. Determine how to parse received data into a usable sql query
//       4. Decide on how to structure the return format (Probably JSON?) Idk if objects can be returned, but if they can it would still limit what this api is compatible with

func main() {
	const (
		server   = "192.168.100.151"
		port     = 1433
		user     = "GoSQLUser"
		password = "terriblepassword123"
	)

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d",
		server, user, password, port)
	database, err := dbhelpers.NewDb(connectionString)
	//This is temporary lol
	if err != nil {
		panic(err)
	}

	container := controllers.Container{Db: database, Tables: retrieveTables(database)}

	tester(&container)
}

func tester(container *controllers.Container) {

	fmt.Println("\nTesting Select all with a valid table")
	//Testing select all
	var testTables = make([]TestTable, 0)

	rows, _ := container.SelectAll("TestTable1")
	for rows.Next() {
		var row = new(TestTable)

		rows.Scan(&row.UserName, &row.Password)
		testTables = append(testTables, *row)
	}
	fmt.Println(testTables)

	//Testing Select all invalid table
	fmt.Println("\nTesting Select all with an invalid table")
	_, err := container.SelectAll("notarealtable")
	fmt.Println(err)
}

//Determine list of valid table names. This is to allow for dynamic queries without having to bind table names to a prepared query, which is unsupported for select statements
//Cheeky workaround lol
func retrieveTables(connection *sql.DB) (tables []string) {
	statement := "SELECT name FROM Sys.Tables"

	rows, _ := connection.Query(statement)

	for rows.Next() {
		var table = new(string)

		rows.Scan(&table)
		tables = append(tables, *table)
	}
	fmt.Println(tables)
	return tables
}

//Temporary testing tables

//TestTable temporary test table
type TestTable struct {
	UserName string
	Password string
}
