package daos

import (
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/models"
)

// NewsDAO persists news data in database
type NewsDAO struct{}

// NewNewsDAO creates a new NewsDAO
func NewNewsDAO() *NewsDAO {
	return &NewsDAO{}
}

// Get reads the news with the specified ID from the database.
func (dao *NewsDAO) Get(rs app.RequestScope, id int) (*models.News, error) {
	var news models.News
	err := rs.Tx().Select().Model(id, &news)
	return &news, err
}

// Create saves a new news record in the database.
// The News.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *NewsDAO) Create(rs app.RequestScope, news *models.News) error {
	news.Id = 0
	return rs.Tx().Model(news).Insert()
}

// Update saves the changes to an news in the database.
func (dao *NewsDAO) Update(rs app.RequestScope, id int, news *models.News) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	news.Id = id
	return rs.Tx().Model(news).Exclude("Id").Update()
}

// Delete deletes an news with the specified ID from the database.
func (dao *NewsDAO) Delete(rs app.RequestScope, id int) error {
	news, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(news).Delete()
}

// Count returns the number of the news records in the database.
func (dao *NewsDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("news").Row(&count)
	return count, err
}

// Query retrieves the news records with the specified offset and limit from the database.
func (dao *NewsDAO) Query(rs app.RequestScope, offset, limit int) ([]models.News, error) {
	news := []models.News{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&news)
	return news, err
}
