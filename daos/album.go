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
func (dao *AlbumDAO) Get(rs app.RequestScope, id int) (*models.Album, error) {
	var album models.Album
	err := rs.Tx().Select().Model(id, &album)
	return &album, err
}

// Get reads the album with the specified ID from the database.
func (dao *AlbumDAO) GetForClient(rs app.RequestScope, id int) (album *models.AlbumForClient, err error) {
	q := rs.Tx().Select("album.id", "title", "artist_id", "artist.name AS artist_name", "album.image").
		From("album").Where(dbx.HashExp{"album.id": id}).
		LeftJoin("artist", dbx.NewExp("\"artist\".\"id\" = \"album\".\"artist_id\""))

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
	album.BeforeInsert()
	res := rs.Tx().Model(album).Insert()
	//rs.Infof("res: %v", res) // res = nil
	//rs.Infof("album: %v", album) // album - mapped row
	album.AfterSave()
	return res
}

// Update saves the changes to an album in the database.
func (dao *AlbumDAO) Update(rs app.RequestScope, id int, album *models.Album) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	album.Id = id
	album.BeforeUpdate()
	res := rs.Tx().Model(album).Exclude("Id").Update()
	album.AfterSave()
	return res
}

// Delete deletes an album with the specified ID from the database.
func (dao *AlbumDAO) Delete(rs app.RequestScope, id int) error {
	album, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	album.BeforeDelete()
	return rs.Tx().Model(album).Delete()
}

// Count returns the number of the album records in the database.
func (dao *AlbumDAO) Count(rs app.RequestScope) (count int, err error) {
	err = rs.Tx().Select("COUNT(*)").From("album").Row(&count)
	return
}

// Query retrieves the album records with the specified offset and limit from the database.
func (dao *AlbumDAO) Query2(rs app.RequestScope, offset, limit int) ([]models.Album, error) {
	albums := []models.Album{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&albums)
	return albums, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (dao *AlbumDAO) Query(rs app.RequestScope, offset, limit int) (albums []models.AlbumForClient, err error) {
	err = rs.Tx().Select("album.id", "title", "artist_id", "artist.name AS artist_name", "album.image").
		From("album").
		LeftJoin("artist", dbx.NewExp("\"artist\".\"id\" = \"album\".\"artist_id\"")).
		OrderBy("album.id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&albums)
	return
}
