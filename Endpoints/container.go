package endpoints

import "database/sql"

//Container Environment container struct used for dependency injection
//Any general utility such as a logger should have a field
type Container struct {
	Db *sql.DB
}

//GetAllTables returns all tables given a map of tables
// func (env Container) GetAllTables() (names []string) {
// 	for key, _ := range env.Tables {
// 		names = append(names, key)
// 	}
// 	return names
// }
