package builders

import (
	"PluckyAPI/Models"
)

//Builder struct contains all required functions for parsing a request into an SQL statement
type Builder struct {
	Tables map[string][]models.Table
}
