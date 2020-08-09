package builders

import (
	"PluckyAPI/Models"
	"fmt"
)

//Insert Handles an insert statement
func (bld SQLBuilder) Insert(request models.Request, query models.Query) (queryWithInsert models.Query, err error) {
	table := request.TableName
	parameters := request.Values
	//Check to see if table exists in db before proceeding
	if !bld.findTable(table) {
		return models.Query{}, fmt.Errorf("table does not exist in the database")
	}

	//Check if columns exist in db before proceeding
	var columns = make([]string, 0)

	for key := range parameters {
		columns = append(columns, key)
	}
	if !bld.findColumns(table, columns) {
		return models.Query{}, fmt.Errorf("Invalid fields")
	}

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
