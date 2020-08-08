package builders

import (
	"PluckyAPI/Models"
	"PluckyAPI/Utilities/misc"
)

//Builder struct contains all required functions for parsing a request into an SQL statement
type Builder struct {
	Tables map[string]models.Table
}

func (bld Builder) findTable(table string) (check bool) {
	var tableNames = make([]string, 0)
	for key := range bld.Tables {
		tableNames = append(tableNames, key)
	}

	return misc.Contains(tableNames, table)
}
