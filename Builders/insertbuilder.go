package builders

import (
	"PluckyAPI/Models"
	"fmt"
)

//Insert Handles an insert statement
func (bld Builder) Insert(table string, parameters map[string]string, query models.Query) (queryWithInsert models.Query, err error) {
	//Check to see if table exists in db before proceeding
	if !bld.findTable(table) {
		return models.Query{}, fmt.Errorf("table does not exist in the database")
	}

	//Check if columns exist in db before proceeding
	var columns = make([]string, 0)

	for key := range parameters {
		columns = append(columns, key)
	}
	//Error is happening here
	if !bld.findColumns(table, columns) {
		return models.Query{}, fmt.Errorf("Invalid fields")
	}

	//Check to see if columns exist in db before proceeding

	//Handling the first part of the SQL statement.

	//Handle the columns in the braces, This can probably be refactored later on to get rid of unnecessary logic - templates?
	var columnString string
	var variableString string
	for index, column := range columns {
		columnString += column
		variableString += fmt.Sprintf("@%v", index)
		if index != len(columns)-1 {
			columnString += ", "
			variableString += ", "
		}
	}

	insertStatement := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", table, columnString, variableString)
	query.AppendToQuery(insertStatement)
	return query, nil
}
