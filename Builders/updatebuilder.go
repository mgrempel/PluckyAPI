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
	counter := 1
	for key, value := range parameters {
		if counter != 1 {
			columnParameter += ", "
		}
		//This isn't accounting for quotations or anything like that, but it's a start
		columnParameter += fmt.Sprintf("%s = %s", key, fmt.Sprintf("@p%v", counter))

		query.AddParameter(value)
		counter++
	}

	updateQuery := fmt.Sprintf("UPDATE %s SET %s", table, columnParameter)

	query.AppendToQuery(updateQuery)
	return query, nil
}
