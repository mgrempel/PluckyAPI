package builders

import (
	"PluckyAPI/Models"
	"fmt"
)

//Where Handles creating the where clause of an sql query. Only handles equals operations for now,
func (bld SQLBuilder) Where(request models.Request, query models.Query) (queryWithWhere models.Query, err error) {

	parameters := request.Values
	whereQueryString := " WHERE "

	fmt.Println(query.GetParams())

	counter := 0
	offset := len(query.GetParams())
	for key, value := range parameters {
		if value != "" {
			whereQueryString += fmt.Sprintf("%s = @%v", key, counter+offset)
			query.AddParameter(value)

			if counter < len(parameters)-1 {
				whereQueryString += " AND "
			}
			counter++
		}
	}

	query.AppendToQuery(whereQueryString)
	return query, nil
}
