package builders

import "fmt"

//Where Handles creating the where clause of an sql query
func (bld Builder) Where(columns []string, values []string) (whereClause string, err error) {
	//Check to ensure each constraint has a matching parameters
	if len(columns) != len(values) {
		return "", fmt.Errorf("Not every constraint has a corresponding value")
	}

	whereClause = " WHERE"

	//Handle first clause

	return whereClause, nil
}
