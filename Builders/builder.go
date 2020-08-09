package builders

import (
	"PluckyAPI/Models"
	"PluckyAPI/Utilities/misc"
)

//Builder struct contains all required functions for parsing a request into an SQL statement
type SQLBuilder struct {
	Tables map[string]models.Table
}

func (bld SQLBuilder) findTable(table string) (check bool) {
	var tableNames = make([]string, 0)
	for key := range bld.Tables {
		tableNames = append(tableNames, key)
	}

	return misc.Contains(tableNames, table)
}

func (bld SQLBuilder) findColumns(tableName string, columns []string) bool {
	tableColumns := bld.Tables[tableName].Columns

	for _, columnName := range tableColumns {
		if !misc.Contains(columns, columnName.ColumnName) {
			return false
		}
	}
	return true
}
