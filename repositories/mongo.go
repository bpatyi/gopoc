package repositories

import "github.com/bpatyi/gopoc/models"

type mongoHandler struct {
}

func NewMongoHandler() Todo {
	return &mongoHandler{}
}

func (m mongoHandler) All() ([]models.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (m mongoHandler) Get(id string) (*models.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (m mongoHandler) Create(todo *models.Todo) error {
	//TODO implement me
	panic("implement me")
}

func (m mongoHandler) Update(id string, todo *models.Todo) {
	//TODO implement me
	panic("implement me")
}

func (m mongoHandler) UpdatePartial(id string, todo *models.Todo) {
	//TODO implement me
	panic("implement me")
}

func (m mongoHandler) Delete(id string) {
	//TODO implement me
	panic("implement me")
}
