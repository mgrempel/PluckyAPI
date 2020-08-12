package endpoints

import (
	interfaces "PluckyAPI/Interfaces"
	models "PluckyAPI/Models"
	"PluckyAPI/Utilities/misc"
	"database/sql"
)

//Container Environment container struct used for dependency injection
//Any general utility such as a logger should have a field
type Container struct {
	Tables  map[string]models.Table
	Db      *sql.DB
	Builder interfaces.Builder
}

func (bld Container) findTable(table string) (check bool) {
	var tableNames = make([]string, 0)
	for key := range bld.Tables {
		tableNames = append(tableNames, key)
	}

	return misc.Contains(tableNames, table)
}

func (bld Container) findColumns(tableName string, columns []string) bool {
	tableColumns := bld.Tables[tableName].Columns

	for _, columnName := range tableColumns {
		if !misc.Contains(columns, columnName.ColumnName) {
			return false
		}
	}
	return true
}
