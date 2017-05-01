package services

import (
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/models"
)

// albumDAO specifies the interface of the album DAO needed by AlbumService.
type albumDAO interface {
	// Get returns the album with the specified the album ID.
	Get(rs app.RequestScope, id int) (*models.Album, error)
	// Count returns the number of albums.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of the albums with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.AlbumForClient, error)
	// Create saves a new album in the storage.
	Create(rs app.RequestScope, album *models.Album) error
	// Update updates the album with the given ID in the storage.
	Update(rs app.RequestScope, id int, album *models.Album) error
	// Delete removes the album with the given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// AlbumService provides services related with albums.
type AlbumService struct {
	dao albumDAO
}

// NewAlbumService creates a new AlbumService with the given album DAO.
func NewAlbumService(dao albumDAO) *AlbumService {
	return &AlbumService{dao}
}

// Get returns the album with the specified the album ID.
func (s *AlbumService) Get(rs app.RequestScope, id int) (*models.Album, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new album.
func (s *AlbumService) Create(rs app.RequestScope, model *models.Album) (*models.Album, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

// Update updates the album with the specified ID.
func (s *AlbumService) Update(rs app.RequestScope, id int, model *models.Album) (*models.Album, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the album with the specified ID.
func (s *AlbumService) Delete(rs app.RequestScope, id int) (*models.Album, error) {
	album, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return album, err
}

// Count returns the number of albums.
func (s *AlbumService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the albums with the specified offset and limit.
func (s *AlbumService) Query(rs app.RequestScope, offset, limit int) ([]models.AlbumForClient, error) {
	return s.dao.Query(rs, offset, limit)
}
