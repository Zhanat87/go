package models

import "github.com/jinzhu/gorm"

type AlbumGorm struct {
	gorm.Model
	title   string
	ArtistID uint
}
