package builders

import (
	"PluckyAPI/Models"
	"fmt"
)

//Select Selects all data from a specified table in the database. Only handling select all for now.
func (bld SQLBuilder) Select(request models.Request, query models.Query) (queryWithSelect models.Query, err error) {
	table := request.TableName

	//No params added here because SQL server is a little oaf.
	selectStatement := fmt.Sprintf("SELECT * FROM %s", table)
	query.AppendToQuery(selectStatement)

	return query, nil
}
