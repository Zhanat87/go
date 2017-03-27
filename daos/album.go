package daos

import (
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/models"
	"github.com/go-ozzo/ozzo-dbx"
)

// AlbumDAO persists album data in database
type AlbumDAO struct{}

// NewAlbumDAO creates a new AlbumDAO
func NewAlbumDAO() *AlbumDAO {
	return &AlbumDAO{}
}

// Get reads the album with the specified ID from the database.
func (dao *AlbumDAO) Get2(rs app.RequestScope, id int) (*models.Album, error) {
	var album models.Album
	err := rs.Tx().Select().Model(id, &album)
	return &album, err
}

// Get reads the album with the specified ID from the database.
func (dao *AlbumDAO) Get(rs app.RequestScope, id int) (album *models.Album, err error) {
	q := rs.Tx().Select("album.id", "title", "artist.name AS artist_name").
		From("album").Where(dbx.HashExp{"album.id": 100}).
		Join("LEFT INNER JOIN", "artist", dbx.NewExp("`artist`.`id` = `album`.`artist_id`"))

	err = q.One(&album)
	if err != nil {
		println("Exec err:", err.Error())
	}
	return
}

// Create saves a new album record in the database.
// The Album.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *AlbumDAO) Create(rs app.RequestScope, album *models.Album) error {
	album.Id = 0
	return rs.Tx().Model(album).Insert()
}

// Update saves the changes to an album in the database.
func (dao *AlbumDAO) Update(rs app.RequestScope, id int, album *models.Album) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	album.Id = id
	return rs.Tx().Model(album).Exclude("Id").Update()
}

// Delete deletes an album with the specified ID from the database.
func (dao *AlbumDAO) Delete(rs app.RequestScope, id int) error {
	album, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(album).Delete()
}

// Count returns the number of the album records in the database.
func (dao *AlbumDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("album").Row(&count)
	return count, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (dao *AlbumDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Album, error) {
	albums := []models.Album{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&albums)
	return albums, err
}
