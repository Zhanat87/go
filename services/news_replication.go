package services

import (
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/models"
)

// newsReplicationDAO specifies the interface of the newsReplication DAO needed by NewsReplicationService.
type newsReplicationDAO interface {
	// Get returns the newsReplication with the specified the newsReplication ID.
	Get(rs app.RequestScope, id int) (*models.NewsReplication, error)
	// Count returns the number of newsReplication.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of the newsReplication with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.NewsReplication, error)
	// Create saves a new newsReplication in the storage.
	Create(rs app.RequestScope, newsReplication *models.NewsReplication) error
	// Update updates the newsReplication with the given ID in the storage.
	Update(rs app.RequestScope, id int, newsReplication *models.NewsReplication) error
	// Delete removes the newsReplication with the given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// NewsReplicationService provides services related with newsReplication.
type NewsReplicationService struct {
	dao newsReplicationDAO
}

// NewNewsReplicationService creates a new NewsReplicationService with the given newsReplication DAO.
func NewNewsReplicationService(dao newsReplicationDAO) *NewsReplicationService {
	return &NewsReplicationService{dao}
}

// Get returns the newsReplication with the specified the newsReplication ID.
func (s *NewsReplicationService) Get(rs app.RequestScope, id int) (*models.NewsReplication, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new newsReplication.
func (s *NewsReplicationService) Create(rs app.RequestScope, model *models.NewsReplication) (*models.NewsReplication, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

// Update updates the newsReplication with the specified ID.
func (s *NewsReplicationService) Update(rs app.RequestScope, id int, model *models.NewsReplication) (*models.NewsReplication, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the newsReplication with the specified ID.
func (s *NewsReplicationService) Delete(rs app.RequestScope, id int) (*models.NewsReplication, error) {
	newsReplication, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return newsReplication, err
}

// Count returns the number of newsReplication.
func (s *NewsReplicationService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the newsReplication with the specified offset and limit.
func (s *NewsReplicationService) Query(rs app.RequestScope, offset, limit int) ([]models.NewsReplication, error) {
	return s.dao.Query(rs, offset, limit)
}
