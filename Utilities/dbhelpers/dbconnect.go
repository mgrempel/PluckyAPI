package dbhelpers

import (
	"database/sql"
	//SQL driver
	_ "github.com/denisenkom/go-mssqldb"
)

//NewDb Instantiates a new connection pool and returns a pointer.
func NewDb(connectionString string) (connection *sql.DB, err error) {
	connection, err = sql.Open("sqlserver", connectionString)
	if err != nil {
		return nil, err
	}
	return connection, nil
}
