package main

import (
	"PluckyAPI/Builders"
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

	tableMap, err := retrieveTables(database)
	fmt.Println(err)

	//Initialize the container for the SQL Builders
	container := builders.Container{Tables: tableMap}
}

//Determine list of valid table names. This is to allow for dynamic queries without having to bind table names to a prepared query, which is unsupported for select statements
//Cheeky workaround lol
func retrieveTables(connection *sql.DB) (tables map[string]builders.Table, err error) {
	tables = make(map[string]builders.Table)

	tableStatement := "SELECT name FROM Sys.Tables"
	rows, err := connection.Query(tableStatement)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var table string

		rows.Scan(&table)

		//Get columns for each table
		columnStatement := fmt.Sprintf("SELECT Column_Name FROM INFORMATION_SCHEMA.COLUMNS WHERE Table_Name = '%s'", table)
		columns, err := connection.Query(columnStatement)
		defer columns.Close()
		if err != nil {
			return nil, err
		}

		var tableColumns = make([]string, 0)

		for columns.Next() {
			var currentColumn string
			columns.Scan(&currentColumn)
			tableColumns = append(tableColumns, currentColumn)
		}
		//add columns into the map
		tables[table] = builders.Table{Columns: tableColumns}
	}
	fmt.Println(tables)
	return tables, nil
}

//Temporary testing tables

//TestTable temporary test table
type TestTable struct {
	UserName string
	Password string
}
