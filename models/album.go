package models

import (
	"github.com/go-ozzo/ozzo-validation"
)

/*
https://github.com/go-ozzo/ozzo-dbx/issues/23
dbx is not orm and can do relational queries, etc with relations
need select by separate queries,
find parent model and then related models by foreign id(s)
 */
type Album struct {
	Id         int    `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	ArtistId   uint   `json:"artistId,string" db:"artist_id"`
	ArtistName string `json:"artistName" db:"artist_name"`
}

// Validate validates the Album fields.
func (m Album) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required, validation.Length(4, 120)),
		validation.Field(&m.ArtistId, validation.Required),
	)
}
