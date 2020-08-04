package controllers

import (
	"PluckyAPI/Utilities/misc"
	"database/sql"
	"fmt"
)

//SelectAll Selects all data from a specified table in the database
func (container Container) SelectAll(table string) (result *sql.Rows, err error) {
	//Check to see if table exists in db before proceeding
	if !misc.Contains(container.GetAllTables(), table) {
		err = fmt.Errorf("No such table \"%s\"", table)
		return nil, err
	}

	db := container.Db
	//Check if db is alive before doing anything
	if err = db.Ping(); err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT * FROM %s;", table)

	//Prepare the Query
	preparedQuery, err := db.Prepare(query)
	defer preparedQuery.Close()
	if err != nil {
		return nil, err
	}

	//Execute query
	result, err = preparedQuery.Query()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// func (container Container) InsertRow() {
//
// }
