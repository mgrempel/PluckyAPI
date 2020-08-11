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
	for key, value := range parameters {
		if _, ok := parameters[key]; ok {
			whereQueryString += fmt.Sprintf("%s = @%v", key, counter)
			query.AddParameter(value)

			if counter < len(parameters)-1 {
				whereQueryString += ", "
			}
			counter++
		}
	}

	query.AppendToQuery(whereQueryString)
	return query, nil
}
