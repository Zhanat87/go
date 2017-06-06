package daos

import (
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/models"
)

// NewsShardDAO persists newsShard data in database
type NewsShardDAO struct{}

// NewNewsShardDAO creates a new NewsShardDAO
func NewNewsShardDAO() *NewsShardDAO {
	return &NewsShardDAO{}
}

// Get reads the newsShard with the specified ID from the database.
func (dao *NewsShardDAO) Get(rs app.RequestScope, id int) (*models.NewsShard, error) {
	var newsShard models.NewsShard
	err := rs.Tx().Select().Model(id, &newsShard)
	return &newsShard, err
}

// Create saves a new newsShard record in the database.
// The NewsShard.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *NewsShardDAO) Create(rs app.RequestScope, newsShard *models.NewsShard) error {
	newsShard.Id = 0
	return rs.Tx().Model(newsShard).Insert()
}

// Update saves the changes to an newsShard in the database.
func (dao *NewsShardDAO) Update(rs app.RequestScope, id int, newsShard *models.NewsShard) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	newsShard.Id = id
	return rs.Tx().Model(newsShard).Exclude("Id").Update()
}

// Delete deletes an newsShard with the specified ID from the database.
func (dao *NewsShardDAO) Delete(rs app.RequestScope, id int) error {
	newsShard, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(newsShard).Delete()
}

// Count returns the number of the newsShard records in the database.
func (dao *NewsShardDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("news_shard").Row(&count)
	return count, err
}

// Query retrieves the newsShard records with the specified offset and limit from the database.
func (dao *NewsShardDAO) Query(rs app.RequestScope, offset, limit int) ([]models.NewsShard, error) {
	newsShard := []models.NewsShard{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&newsShard)
	return newsShard, err
}
