package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

// Artist represents an artist record.
type ArtistGorm struct {
	gorm.Model
	Name string `json:"name" db:"name"`
}

// Validate validates the Artist fields.
func (m Artist) ValidateArtistGorm() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 120)),
	)
}
