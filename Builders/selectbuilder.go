package builders

import (
	"PluckyAPI/Models"
	"fmt"
)

//Select Selects all data from a specified table in the database. Only handling select all for now.
func (bld Builder) Select(table string, query models.Query) (queryWithSelect models.Query, err error) {
	//Check to see if table exists in db before proceeding
	if bld.Tables[table] == nil {
		return models.Query{}, fmt.Errorf("table does not exist in the database")
	}

	//No params added here because SQL server is a little oaf.
	selectStatement := fmt.Sprintf("SELECT * FROM %s", table)
	query.AppendToQuery(selectStatement)

	return query, nil

	// //Prepare the Query
	// preparedQuery, err := db.Prepare(query)
	// defer preparedQuery.Close()
	// if err != nil {
	// 	return nil, err
	// }
	//
	// //Execute query
	// result, err = preparedQuery.Query()
	// if err != nil {
	// 	return nil, err
	// }
}

//Select with parameters

//func (container Container) ParameterSelect(table string, parameters ...string)

//Insert row

//Delete row
