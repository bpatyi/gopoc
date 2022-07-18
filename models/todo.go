package models

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Todo struct {
	ID          string
	Title       string
	Description string
}

func (t Todo) Validate() error {
	return validation.ValidateStruct(
		&t,
		validation.Field(&t.ID, validation.Required),
		validation.Field(&t.Title, validation.Required),
		validation.Field(&t.Description, validation.Required),
	)
}
