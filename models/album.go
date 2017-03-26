package models

type Album struct {
	id       int    `json:"id" db:"id"`
	title    string  `json:"title" db:"title"`
	artistId uint    `json:"artistId" db:"artist_id"`
}
