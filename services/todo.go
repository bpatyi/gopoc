package services

import (
	"github.com/bpatyi/gopoc/models"
	"github.com/bpatyi/gopoc/repositories"
)

type Todo interface {
	GetAll() ([]models.Todo, error)
	Create(todo *models.Todo) (*models.Todo, error)
	GetByID(id string) (*models.Todo, error)
	Update(id string, todo *models.Todo) (*models.Todo, error)
	UpdatePartial(id string, todo *models.Todo) (*models.Todo, error)
	Delete(id string) error
}

type todoService struct {
	repository repositories.Todo
}

func NewTodoService(repository repositories.Todo) Todo {
	return &todoService{
		repository: repository,
	}
}

func (t *todoService) GetAll() ([]models.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t *todoService) Create(todo *models.Todo) (*models.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t *todoService) GetByID(id string) (*models.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t *todoService) Update(id string, todo *models.Todo) (*models.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t *todoService) UpdatePartial(id string, todo *models.Todo) (*models.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t *todoService) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
