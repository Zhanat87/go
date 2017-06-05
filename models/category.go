package models

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Category struct {
	Id         int    `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
}

// Validate validates the Category fields.
func (m Category) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required, validation.Length(4, 120)),
	)
}
