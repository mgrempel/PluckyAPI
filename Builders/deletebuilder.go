package builders

import (
	models "PluckyAPI/Models"
	"fmt"
)

//Delete prepares the first portion of a delete statement for the database
func (bld SQLBuilder) Delete(request models.Request, query models.Query) (queryWithSelect models.Query, err error) {
	table := request.TableName

	//No params added here because SQL server is a little oaf.
	deleteStatement := fmt.Sprintf("DELETE FROM %s", table)
	query.AppendToQuery(deleteStatement)

	return query, nil
}
