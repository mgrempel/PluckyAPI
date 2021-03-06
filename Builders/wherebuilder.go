package builders

import (
	"PluckyAPI/Models"
	"fmt"
)

//Where Handles creating the where clause of an sql query. Only handles equals operations for now,
func (bld SQLBuilder) Where(request models.Request, query models.Query) (queryWithWhere models.Query, err error) {

	parameters := request.Values
	whereQueryString := " WHERE "

	counter := 0
	offset := len(query.GetParams()) + 1
	for key, value := range parameters {
		if value != "" {
			if counter != 0 {
				whereQueryString += " AND "
			}

			whereQueryString += fmt.Sprintf("%s = @p%v", key, counter+offset)
			query.AddParameter(value)
			counter++
		}
	}

	query.AppendToQuery(whereQueryString)
	return query, nil
}
