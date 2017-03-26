package models

type Album struct {
	Id       int    `json:"id" db:"id"`
	Title    string  `json:"title" db:"title"`
	ArtistId uint    `json:"artistId" db:"artist_id"`
}
