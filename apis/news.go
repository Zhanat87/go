package apis

import (
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/models"
)

type (
	// newsService specifies the interface for the news service needed by newsResource.
	newsService interface {
		Get(rs app.RequestScope, id int) (*models.News, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.News, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.News) (*models.News, error)
		Update(rs app.RequestScope, id int, model *models.News) (*models.News, error)
		Delete(rs app.RequestScope, id int) (*models.News, error)
	}

	// newsResource defines the handlers for the CRUD APIs.
	newsResource struct {
		service newsService
	}
)

// ServeNews sets up the routing of news endpoints and the corresponding handlers.
func ServeNewsResource(rg *routing.RouteGroup, service newsService) {
	r := &newsResource{service}
	rg.Get("/news/<id>", r.get)
	rg.Get("/news", r.query)
	rg.Post("/news", r.create)
	rg.Put("/news/<id>", r.update)
	rg.Delete("/news/<id>", r.delete)
}

func (r *newsResource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *newsResource) query(c *routing.Context) error {
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

func (r *newsResource) create(c *routing.Context) error {
	var model models.News
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *newsResource) update(c *routing.Context) error {
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

func (r *newsResource) delete(c *routing.Context) error {
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
