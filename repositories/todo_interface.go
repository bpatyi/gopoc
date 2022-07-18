package repositories

import "github.com/bpatyi/gopoc/models"

type TodoRepository interface {
	All() ([]models.Todo, error)
	Get(id string) (*models.Todo, error)
	Create(todo *models.Todo) error
	Update(id string, todo *models.Todo)
	Delete(id string)
}
