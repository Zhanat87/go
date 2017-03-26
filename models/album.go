package models

import "github.com/go-ozzo/ozzo-validation"

type Album struct {
	Id       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	ArtistId uint   `json:"artistId" db:"artist_id"`
}

// Validate validates the Artist fields.
func (m Album) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required, validation.Length(0, 120)),
		validation.Field(&m.ArtistId, validation.Required),
	)
}
