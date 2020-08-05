package models

//Table struct representing a database table
type Table struct {
	Columns []Column
}

//Column struct representing a database column
type Column struct {
	ColumnName string
	ColumnType string
}
