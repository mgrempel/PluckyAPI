package builders

import (
	"PluckyAPI/Models"
	"fmt"
)

//Insert Handles an insert statement
func (bld SQLBuilder) Insert(request models.Request, query models.Query) (queryWithInsert models.Query, err error) {
	table := request.TableName
	parameters := request.Values

	//Handle the columns in the braces, This can probably be refactored later on to get rid of unnecessary logic - templates?
	var columns = make([]string, 0)
	for key := range parameters {
		columns = append(columns, key)
	}

	var columnString string
	var variableString string
	for index, column := range columns {
		columnString += column
		variableString += fmt.Sprintf("@p%v", index)
		query.AddParameter(column)
		if index != len(columns)-1 {
			columnString += ", "
			variableString += ", "
		}
	}

	insertStatement := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", table, columnString, variableString)
	query.AppendToQuery(insertStatement)
	return query, nil
}
