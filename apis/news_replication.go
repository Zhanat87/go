package apis

import (
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/models"
	"github.com/go-ozzo/ozzo-dbx"
)

type (
	// newsReplicationService specifies the interface for the newsReplication service needed by newsReplicationResource.
	newsReplicationService interface {
		Get(rs app.RequestScope, id int) (*models.NewsReplication, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.NewsReplication, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.NewsReplication) (*models.NewsReplication, error)
		Update(rs app.RequestScope, id int, model *models.NewsReplication) (*models.NewsReplication, error)
		Delete(rs app.RequestScope, id int) (*models.NewsReplication, error)
	}

	// newsReplicationResource defines the handlers for the CRUD APIs.
	newsReplicationResource struct {
		service newsReplicationService
	}
)

// ServeNewsReplication sets up the routing of newsReplication endpoints and the corresponding handlers.
func ServeNewsReplicationResource(rg *routing.RouteGroup, service newsReplicationService) {
	r := &newsReplicationResource{service}
	rg.Get("/replication/news/<id>", r.get)
	rg.Get("/replication/news", r.query)
	rg.Post("/replication/news", r.create)
	rg.Put("/replication/news/<id>", r.update)
	rg.Delete("/replication/news/<id>", r.delete)
}

func (r *newsReplicationResource) get(c *routing.Context) error {
	setReplicationDbConnection(c, false)
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

func (r *newsReplicationResource) query(c *routing.Context) error {
	setReplicationDbConnection(c, false)
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

func (r *newsReplicationResource) create(c *routing.Context) error {
	setReplicationDbConnection(c, true)
	var model models.NewsReplication
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *newsReplicationResource) update(c *routing.Context) error {
	setReplicationDbConnection(c, true)
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

func (r *newsReplicationResource) delete(c *routing.Context) error {
	setReplicationDbConnection(c, true)
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

func setReplicationDbConnection(c *routing.Context, isMaster bool) {
	var dsn string
	if isMaster {
		dsn = app.Config.DSN_DOCKER_COMPOSE_V3_REPLICATION_MASTER
	} else {
		dsn = app.Config.DSN_DOCKER_COMPOSE_V3_REPLICATION_SLAVE
	}
	db, err := dbx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	app.GetRequestScope(c).SetTx(db)
}