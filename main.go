package main

import (
	builders "PluckyAPI/Builders"
	endpoints "PluckyAPI/Endpoints"
	"PluckyAPI/Utilities/dbhelpers"
	"fmt"
	"net/http"
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

	tableMap, err := dbhelpers.PopulateTables(database)
	if err != nil {
		panic(err)
	}

	builder := builders.SQLBuilder{}
	container := endpoints.Container{Db: database, Builder: builder, Tables: tableMap}

	http.HandleFunc("/api", container.HandleRequest)
	http.ListenAndServe(":8080", nil)
}
