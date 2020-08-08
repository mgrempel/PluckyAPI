package endpoints

import (
	"PluckyAPI/Builders"
	"database/sql"
)

//Container Environment container struct used for dependency injection
//Any general utility such as a logger should have a field
type Container struct {
	Db      *sql.DB
	Builder builders.Builder
}
