package services

import (
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/models"
)

// newsShardDAO specifies the interface of the newsShard DAO needed by NewsShardService.
type newsShardDAO interface {
	// Get returns the newsShard with the specified the newsShard ID.
	Get(rs app.RequestScope, id int) (*models.NewsShard, error)
	// Count returns the number of newsShard.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of the newsShard with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.NewsShard, error)
	// Create saves a new newsShard in the storage.
	Create(rs app.RequestScope, newsShard *models.NewsShard) error
	// Update updates the newsShard with the given ID in the storage.
	Update(rs app.RequestScope, id int, newsShard *models.NewsShard) error
	// Delete removes the newsShard with the given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// NewsShardService provides services related with newsShard.
type NewsShardService struct {
	dao newsShardDAO
}

// NewNewsShardService creates a new NewsShardService with the given newsShard DAO.
func NewNewsShardService(dao newsShardDAO) *NewsShardService {
	return &NewsShardService{dao}
}

// Get returns the newsShard with the specified the newsShard ID.
func (s *NewsShardService) Get(rs app.RequestScope, id int) (*models.NewsShard, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new newsShard.
func (s *NewsShardService) Create(rs app.RequestScope, model *models.NewsShard) (*models.NewsShard, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

// Update updates the newsShard with the specified ID.
func (s *NewsShardService) Update(rs app.RequestScope, id int, model *models.NewsShard) (*models.NewsShard, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the newsShard with the specified ID.
func (s *NewsShardService) Delete(rs app.RequestScope, id int) (*models.NewsShard, error) {
	newsShard, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return newsShard, err
}

// Count returns the number of newsShard.
func (s *NewsShardService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the newsShard with the specified offset and limit.
func (s *NewsShardService) Query(rs app.RequestScope, offset, limit int) ([]models.NewsShard, error) {
	return s.dao.Query(rs, offset, limit)
}
