package apis

import (
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/models"
)

type (
	// albumService specifies the interface for the album service needed by albumResource.
	albumService interface {
		Get(rs app.RequestScope, id int) (*models.Album, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Album, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Album) (*models.Album, error)
		Update(rs app.RequestScope, id int, model *models.Album) (*models.Album, error)
		Delete(rs app.RequestScope, id int) (*models.Album, error)
	}

	// albumResource defines the handlers for the CRUD APIs.
	albumResource struct {
		service albumService
	}
)

// ServeAlbum sets up the routing of album endpoints and the corresponding handlers.
func ServeAlbumResource(rg *routing.RouteGroup, service albumService) {
	r := &albumResource{service}
	rg.Get("/albums/<id>", r.get)
	rg.Get("/albums", r.query)
	rg.Post("/albums", r.create)
	rg.Put("/albums/<id>", r.update)
	rg.Delete("/albums/<id>", r.delete)
}

func (r *albumResource) get(c *routing.Context) (err error) {
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

func (r *albumResource) query(c *routing.Context) error {
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

func (r *albumResource) create(c *routing.Context) error {
	var model models.Album
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *albumResource) update(c *routing.Context) error {
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

func (r *albumResource) delete(c *routing.Context) error {
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
