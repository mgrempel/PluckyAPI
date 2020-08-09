package interfaces

import (
	"PluckyAPI/Models"
)

//Builder interface for creating queries to be run against the database
type Builder interface {
	Insert(models.Request, models.Query) models.Query
	Select(models.Request, models.Query) models.Query
	Update(models.Request, models.Query) models.Query
	Delete(models.Request, models.Query) models.Query
	Where(models.Request, models.Query) models.Query
}
