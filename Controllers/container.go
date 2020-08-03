package controllers

import "database/sql"

//Container Environment container struct used for dependency injection
//Any general utility such as a logger should have a field
type Container struct {
	Db     *sql.DB
	Tables []string
}
