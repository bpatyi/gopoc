package repositories

import "github.com/bpatyi/gopoc/models"

type Todo interface {
	All() ([]models.Todo, error)
	Get(id string) (*models.Todo, error)
	Create(todo *models.Todo) error
	Update(id string, todo *models.Todo)
	UpdatePartial(id string, todo *models.Todo)
	Delete(id string)
}
