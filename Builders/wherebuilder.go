package builders

import (
	"PluckyAPI/Models"
)

//Where Handles creating the where clause of an sql query
func (bld SQLBuilder) Where(request models.Request, query models.Query) (queryWithWhere models.Query, err error) {
	//Check to ensure each constraint has a matching parameters

	query.AppendToQuery(" WHERE @1")

	return queryWithWhere, nil
}
