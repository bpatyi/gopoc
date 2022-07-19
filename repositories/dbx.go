package repositories

import "github.com/bpatyi/gopoc/models"

type dbxHandler struct{}

func NewDbxHandler() Todo {
	return &dbxHandler{}
}

func (d dbxHandler) All() ([]models.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbxHandler) Get(id string) (*models.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbxHandler) Create(todo *models.Todo) error {
	//TODO implement me
	panic("implement me")
}

func (d dbxHandler) Update(id string, todo *models.Todo) {
	//TODO implement me
	panic("implement me")
}

func (d dbxHandler) UpdatePartial(id string, todo *models.Todo) {
	//TODO implement me
	panic("implement me")
}

func (d dbxHandler) Delete(id string) {
	//TODO implement me
	panic("implement me")
}
