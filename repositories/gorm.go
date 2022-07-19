package repositories

import "github.com/bpatyi/gopoc/models"

type gormHandler struct {
}

func NewGormHandler() Todo {
	return &gormHandler{}
}

func (g gormHandler) All() ([]models.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (g gormHandler) Get(id string) (*models.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (g gormHandler) Create(todo *models.Todo) error {
	//TODO implement me
	panic("implement me")
}

func (g gormHandler) Update(id string, todo *models.Todo) {
	//TODO implement me
	panic("implement me")
}

func (g gormHandler) UpdatePartial(id string, todo *models.Todo) {
	//TODO implement me
	panic("implement me")
}

func (g gormHandler) Delete(id string) {
	//TODO implement me
	panic("implement me")
}
