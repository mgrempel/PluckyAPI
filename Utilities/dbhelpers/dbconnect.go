package dbhelpers

import (
	"database/sql"
	//SQL driver
	_ "github.com/denisenkom/go-mssqldb"
)

//NewDb Instantiates a new connection pool and returns a pointer.
func NewDb(connectionString string) (connection *sql.DB, err error) {
	return sql.Open("sqlserver", connectionString)
}
