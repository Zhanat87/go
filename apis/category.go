package apis

import (
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/models"
)

type (
	// categoryService specifies the interface for the category service needed by categoryResource.
	categoryService interface {
		Get(rs app.RequestScope, id int) (*models.Category, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Category, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Category) (*models.Category, error)
		Update(rs app.RequestScope, id int, model *models.Category) (*models.Category, error)
		Delete(rs app.RequestScope, id int) (*models.Category, error)
	}

	// categoryResource defines the handlers for the CRUD APIs.
	categoryResource struct {
		service categoryService
	}
)

// ServeCategory sets up the routing of category endpoints and the corresponding handlers.
func ServeCategoryResource(rg *routing.RouteGroup, service categoryService) {
	r := &categoryResource{service}
	rg.Get("/categories/<id>", r.get)
	rg.Get("/categories", r.query)
	rg.Post("/categories", r.create)
	rg.Put("/categories/<id>", r.update)
	rg.Delete("/categories/<id>", r.delete)
}

func (r *categoryResource) get(c *routing.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	response, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return
	}

	return c.Write(response)
}

func (r *categoryResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return err
	}
	paginatedList.Items = items
	return c.Write(paginatedList)
}

func (r *categoryResource) create(c *routing.Context) error {
	var model models.Category
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *categoryResource) update(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	rs := app.GetRequestScope(c)

	model, err := r.service.Get(rs, id)
	if err != nil {
		return err
	}

	if err := c.Read(model); err != nil {
		return err
	}

	response, err := r.service.Update(rs, id, model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *categoryResource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Delete(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}
