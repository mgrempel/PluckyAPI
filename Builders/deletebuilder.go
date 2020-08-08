package builders

import (
	"PluckyAPI/Models"
	"fmt"
)

//Delete prepares the first portion of a delete statement for the database
func (bld Builder) Delete(table string, query models.Query) (queryWithSelect models.Query, err error) {
	//Check to see if table exists in db before proceeding
	if !bld.findTable(table) {
		return models.Query{}, fmt.Errorf("table does not exist in the database")
	}

	//No params added here because SQL server is a little oaf.
	deleteStatement := fmt.Sprintf("DELETE FROM %s", table)
	query.AppendToQuery(deleteStatement)

	return query, nil
}