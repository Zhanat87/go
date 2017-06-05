package services

import (
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/models"
)

// newsDAO specifies the interface of the news DAO needed by NewsService.
type newsDAO interface {
	// Get returns the news with the specified the news ID.
	Get(rs app.RequestScope, id int) (*models.News, error)
	// Count returns the number of news.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of the news with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.News, error)
	// Create saves a new news in the storage.
	Create(rs app.RequestScope, news *models.News) error
	// Update updates the news with the given ID in the storage.
	Update(rs app.RequestScope, id int, news *models.News) error
	// Delete removes the news with the given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// NewsService provides services related with news.
type NewsService struct {
	dao newsDAO
}

// NewNewsService creates a new NewsService with the given news DAO.
func NewNewsService(dao newsDAO) *NewsService {
	return &NewsService{dao}
}

// Get returns the news with the specified the news ID.
func (s *NewsService) Get(rs app.RequestScope, id int) (*models.News, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new news.
func (s *NewsService) Create(rs app.RequestScope, model *models.News) (*models.News, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

// Update updates the news with the specified ID.
func (s *NewsService) Update(rs app.RequestScope, id int, model *models.News) (*models.News, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the news with the specified ID.
func (s *NewsService) Delete(rs app.RequestScope, id int) (*models.News, error) {
	news, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return news, err
}

// Count returns the number of news.
func (s *NewsService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the news with the specified offset and limit.
func (s *NewsService) Query(rs app.RequestScope, offset, limit int) ([]models.News, error) {
	return s.dao.Query(rs, offset, limit)
}
