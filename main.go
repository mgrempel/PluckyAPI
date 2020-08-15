package main

import (
	builders "PluckyAPI/Builders"
	endpoints "PluckyAPI/Endpoints"
	"PluckyAPI/Utilities/confighelper"
	"PluckyAPI/Utilities/dbhelpers"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	//Parameters for the connection string, these will eventually be pulled from a config file
	// const (
	// 	server   = "192.168.100.151"
	// 	port     = 1433
	// 	user     = "GoSQLUser"
	// 	password = "terriblepassword123"
	// )

	//Get config object from confighelper
	config, err := confighelper.GetConfigValues()
	if err != nil {
		panic("Invalid config file")
	}

	//Create connection string
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d",
		config.ServerAddress, config.User, config.Password, config.Port)
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

	router := httprouter.New()
	router.POST("/", container.HandleRequest)
	http.ListenAndServe(":8080", router)
}
