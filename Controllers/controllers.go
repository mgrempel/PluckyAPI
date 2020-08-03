package controllers

import "database/sql"

//SelectAll Selects all data from a specified table in the database
func (container *Container) SelectAll(table string) (result *sql.Rows, err error) {
	db := container.Db
	baseQuery := "SELECT * FROM @TableName;"

	//Check if db is alive before doing anything
	if err = db.Ping(); err != nil {
		return result, err
	}

	//Prepare the Query
	preparedQuery, err := db.Prepare(baseQuery)
	defer preparedQuery.Close()

	//Execute query
	result, err = preparedQuery.Query(preparedQuery,
		sql.Named("TableName", table))

	return result, err
}
