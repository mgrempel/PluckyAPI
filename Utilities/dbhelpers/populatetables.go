package dbhelpers

import (
	models "PluckyAPI/Models"
	"database/sql"
)

//PopulateTables - Handles the initial population of table information which is stored in memory, more or less to prevent an sql injection attack
//Since the framework is modelless
func PopulateTables(db *sql.DB) (tables map[string]models.Table, err error) {
	tables = make(map[string]models.Table)
	//Get all table
	tableStatement := "SELECT name FROM Sys.Tables"
	tableNames, err := db.Query(tableStatement)
	defer tableNames.Close()
	if err != nil {
		return nil, err
	}

	rowStatement := "SELECT Column_Name, Data_Type FROM INFORMATION_SCHEMA.COLUMNS WHERE Table_Name = @TableName"
	rowQuery, err := db.Prepare(rowStatement)
	defer rowQuery.Close()
	if err != nil {
		return nil, err
	}

	for tableNames.Next() {
		var table string
		var tableStruct = models.Table{}

		err := tableNames.Scan(&table)
		if err != nil {
			return nil, err
		}
		//Query for rows associated with current table
		rows, err := rowQuery.Query(rowStatement, sql.Named("TableName", table))
		defer rows.Close()
		if err != nil {
			return nil, err
		}
		//Grab information for each column and populate a column struct
		for rows.Next() {
			var column = models.Column{}

			err := rows.Scan(&column.ColumnName, &column.ColumnType)
			if err != nil {
				return nil, err
			}
			tableStruct.Columns = append(tableStruct.Columns, column)
		}
		//All information on columns has been grabbed, insert table into map
		tables[table] = tableStruct
	}
	return tables, nil
}
