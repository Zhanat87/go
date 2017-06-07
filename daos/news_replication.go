package daos

import (
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/models"
)

// NewsReplicationDAO persists newsReplication data in database
type NewsReplicationDAO struct{}

// NewNewsReplicationDAO creates a new NewsReplicationDAO
func NewNewsReplicationDAO() *NewsReplicationDAO {
	return &NewsReplicationDAO{}
}

// Get reads the newsReplication with the specified ID from the database.
func (dao *NewsReplicationDAO) Get(rs app.RequestScope, id int) (*models.NewsReplication, error) {
	var newsReplication models.NewsReplication
	err := rs.Tx().Select().Model(id, &newsReplication)
	return &newsReplication, err
}

// Create saves a new newsReplication record in the database.
// The NewsReplication.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *NewsReplicationDAO) Create(rs app.RequestScope, newsReplication *models.NewsReplication) error {
	newsReplication.Id = 0
	return rs.Tx().Model(newsReplication).Insert()
}

// Update saves the changes to an newsReplication in the database.
func (dao *NewsReplicationDAO) Update(rs app.RequestScope, id int, newsReplication *models.NewsReplication) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	newsReplication.Id = id
	return rs.Tx().Model(newsReplication).Exclude("Id").Update()
}

// Delete deletes an newsReplication with the specified ID from the database.
func (dao *NewsReplicationDAO) Delete(rs app.RequestScope, id int) error {
	newsReplication, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(newsReplication).Delete()
}

// Count returns the number of the newsReplication records in the database.
func (dao *NewsReplicationDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("news_replication").Row(&count)
	return count, err
}

// Query retrieves the newsReplication records with the specified offset and limit from the database.
func (dao *NewsReplicationDAO) Query(rs app.RequestScope, offset, limit int) ([]models.NewsReplication, error) {
	newsReplication := []models.NewsReplication{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&newsReplication)
	return newsReplication, err
}
