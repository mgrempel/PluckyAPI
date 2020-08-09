package builders

import (
	models "PluckyAPI/Models"
	"fmt"
)

//Update Parses a query into an update statement
func (bld SQLBuilder) Update(request models.Request, query models.Query) (queryWithUpdate models.Query, err error) {
	table := request.TableName
	parameters := request.Updates
	//Check if table exists before proceeding
	if !bld.findTable(table) {
		return models.Query{}, fmt.Errorf("table does not exist in the database")
	}
	//Check to see if columns exist before proceeding
	var columns = make([]string, 0)

	for key := range parameters {
		columns = append(columns, key)
	}
	if !bld.findColumns(table, columns) {
		return models.Query{}, fmt.Errorf("Invalid fields")
	}

	var columnParameter string

	//I need to worry about the types of the columns here as well.
	counter := 0
	for key := range parameters {
		//This isn't accounting for quotations or anything like that, but it's a start
		columnParameter += fmt.Sprintf("%s = %s", key, fmt.Sprintf("@%v", counter))

		if counter != len(parameters)-1 {
			columnParameter += ", "
		}
		counter++
	}

	updateQuery := fmt.Sprintf("UPDATE %s SET %s", table, columnParameter)

	queryWithUpdate.AppendToQuery(updateQuery)
	return queryWithUpdate, nil
}
