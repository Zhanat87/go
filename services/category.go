package services

import (
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/models"
)

// categoryDAO specifies the interface of the category DAO needed by CategoryService.
type categoryDAO interface {
	// Get returns the category with the specified the category ID.
	Get(rs app.RequestScope, id int) (*models.Category, error)
	// Count returns the number of categories.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of the categories with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.Category, error)
	// Create saves a new category in the storage.
	Create(rs app.RequestScope, category *models.Category) error
	// Update updates the category with the given ID in the storage.
	Update(rs app.RequestScope, id int, category *models.Category) error
	// Delete removes the category with the given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// CategoryService provides services related with categories.
type CategoryService struct {
	dao categoryDAO
}

// NewCategoryService creates a new CategoryService with the given category DAO.
func NewCategoryService(dao categoryDAO) *CategoryService {
	return &CategoryService{dao}
}

// Get returns the category with the specified the category ID.
func (s *CategoryService) Get(rs app.RequestScope, id int) (*models.Category, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new category.
func (s *CategoryService) Create(rs app.RequestScope, model *models.Category) (*models.Category, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

// Update updates the category with the specified ID.
func (s *CategoryService) Update(rs app.RequestScope, id int, model *models.Category) (*models.Category, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the category with the specified ID.
func (s *CategoryService) Delete(rs app.RequestScope, id int) (*models.Category, error) {
	category, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return category, err
}

// Count returns the number of categories.
func (s *CategoryService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the categories with the specified offset and limit.
func (s *CategoryService) Query(rs app.RequestScope, offset, limit int) ([]models.Category, error) {
	return s.dao.Query(rs, offset, limit)
}
