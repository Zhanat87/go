package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"database/sql"
	"github.com/satori/go.uuid"
	"github.com/Zhanat87/go/helpers"
)

const ARTISTS_BASE_IMAGE_PATH = "static/artists/images/"

// Artist represents an artist record.
type Artist struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	ImageBase64 JsonNullString `json:"image_base_64" db:"-"`
	Image       JsonNullString `json:"image,omitempty" db:"image"`
}

// Validate validates the Artist fields.
func (m Artist) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(4, 120)),
	)
}

func (m *Artist) BeforeInsert() {
	m.SaveImage()
}

func (m *Artist) BeforeUpdate() {
	m.SaveImage()
}

func (m *Artist) BeforeDelete() {
	m.RemoveImage()
}

func (m *Artist) SaveImage() {
	imageNotEmpty := m.Image.Valid
	imageString := m.Image.String
	if m.ImageBase64.Valid {
		if m.ImageBase64.String == "remove" {
			helpers.RemoveImageWithThumbnails(ARTISTS_BASE_IMAGE_PATH, imageString)
			m.Image = JsonNullString{sql.NullString{String:"", Valid:false}}
		} else {
			imageUUID := uuid.NewV4()
			img, err := helpers.SaveImageToDisk(ARTISTS_BASE_IMAGE_PATH, imageUUID.String(), m.ImageBase64.String)
			if err != nil {
				panic(err)
			}
			helpers.MakeThumbnails(img, ARTISTS_BASE_IMAGE_PATH)
			if imageNotEmpty {
				helpers.RemoveImageWithThumbnails(ARTISTS_BASE_IMAGE_PATH, imageString)
			}
			m.Image = JsonNullString{sql.NullString{String:img, Valid:true}}
		}
	}
}

func (m *Artist) RemoveImage() {
	imageNotEmpty := m.Image.Valid
	imageString := m.Image.String
	if imageNotEmpty {
		helpers.RemoveImageWithThumbnails(ARTISTS_BASE_IMAGE_PATH, imageString)
	}
}
