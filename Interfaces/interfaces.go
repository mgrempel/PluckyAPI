package interfaces

import (
	models "PluckyAPI/Models"
)

//Builder interface for creating queries to be run against the database
type Builder interface {
	Insert(models.Request, models.Query) (models.Query, error)
	Select(models.Request, models.Query) (models.Query, error)
	Update(models.Request, models.Query) (models.Query, error)
	Delete(models.Request, models.Query) (models.Query, error)
	Where(models.Request, models.Query) (models.Query, error)
}
