package builders

import (
	"fmt"
)

//Select Selects all data from a specified table in the database. Only handling select all for now.
func (bld Builder) Select(table string) (selectStatment string, err error) {
	//Check to see if table exists in db before proceeding
	if bld.Tables[table] == nil {
		return "", fmt.Errorf("table does not exist in the database")
	}

	return fmt.Sprintf("SELECT * FROM %s", table), nil

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
