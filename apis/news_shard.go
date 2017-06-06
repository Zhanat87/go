package apis

import (
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/models"
)

type (
	// newsShardService specifies the interface for the newsShard service needed by newsShardResource.
	newsShardService interface {
		Get(rs app.RequestScope, id int) (*models.NewsShard, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.NewsShard, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.NewsShard) (*models.NewsShard, error)
		Update(rs app.RequestScope, id int, model *models.NewsShard) (*models.NewsShard, error)
		Delete(rs app.RequestScope, id int) (*models.NewsShard, error)
	}

	// newsShardResource defines the handlers for the CRUD APIs.
	newsShardResource struct {
		service newsShardService
	}
)

// ServeNewsShard sets up the routing of newsShard endpoints and the corresponding handlers.
func ServeNewsShardResource(rg *routing.RouteGroup, service newsShardService) {
	r := &newsShardResource{service}
	rg.Get("/shard/news/<id>", r.get)
	rg.Get("/shard/news", r.query)
	rg.Post("/shard/news", r.create)
	rg.Put("/shard/news/<id>", r.update)
	rg.Delete("/shard/news/<id>", r.delete)
}

func (r *newsShardResource) get(c *routing.Context) error {
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

func (r *newsShardResource) query(c *routing.Context) error {
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

func (r *newsShardResource) create(c *routing.Context) error {
	var model models.NewsShard
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *newsShardResource) update(c *routing.Context) error {
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

func (r *newsShardResource) delete(c *routing.Context) error {
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
