package models

import (
	"github.com/go-ozzo/ozzo-validation"
)

// NewsReplication represents an newsReplication record.
type NewsReplication struct {
	Id         int    `json:"id" db:"id"`
	CategoryId uint   `json:"category_id,string" db:"category_id"`
	Author     string `json:"author" db:"author"`
	Rate       int    `json:"rate,string" db:"rate"`
	Title      string `json:"title" db:"title"`
	Text       string `json:"text" db:"text"`
}

// Validate validates the NewsReplication fields.
func (m NewsReplication) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required, validation.Length(4, 120)),
		validation.Field(&m.CategoryId, validation.Required),
		validation.Field(&m.Rate, validation.Required),
		validation.Field(&m.Text, validation.Required),
		validation.Field(&m.Author, validation.Required),
	)
}
