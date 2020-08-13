package builders

import (
	models "PluckyAPI/Models"
	"fmt"
)

//Update Parses a query into an update statement
func (bld SQLBuilder) Update(request models.Request, query models.Query) (queryWithUpdate models.Query, err error) {
	table := request.TableName
	parameters := request.Updates

	var columnParameter string
	counter := 0
	for key, value := range parameters {
		//This isn't accounting for quotations or anything like that, but it's a start
		columnParameter += fmt.Sprintf("%s = %s", key, fmt.Sprintf("@p%v", counter))

		if counter != len(parameters)-1 {
			columnParameter += ", "
		}
		query.AddParameter(value)
		counter++
	}

	updateQuery := fmt.Sprintf("UPDATE %s SET %s", table, columnParameter)

	query.AppendToQuery(updateQuery)
	return query, nil
}
