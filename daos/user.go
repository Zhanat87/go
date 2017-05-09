package daos

import (
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/models"
	"github.com/go-ozzo/ozzo-dbx"
)

// UserDAO persists user data in database
type UserDAO struct{}

// NewUserDAO creates a new UserDAO
func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

// Get reads the user with the specified ID from the database.
func (dao *UserDAO) Get(rs app.RequestScope, id int) (*models.User, error) {
	var user models.User
	err := rs.Tx().Select().Model(id, &user)
	return &user, err
}

// Create saves a new user record in the database.
// The User.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *UserDAO) Create(rs app.RequestScope, user *models.User) error {
	user.Id = 0
	user.BeforeInsert()
	return rs.Tx().Model(user).Insert()
}

// Update saves the changes to an user in the database.
func (dao *UserDAO) Update(rs app.RequestScope, id int, user *models.User) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	user.Id = id
	user.BeforeUpdate()
	return rs.Tx().Model(user).Exclude("Id").Update()
}

// Delete deletes an user with the specified ID from the database.
func (dao *UserDAO) Delete(rs app.RequestScope, id int) error {
	user, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(user).Delete()
}

// Count returns the number of the user records in the database.
func (dao *UserDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("user").Row(&count)
	return count, err
}

// Query retrieves the user records with the specified offset and limit from the database.
func (dao *UserDAO) Query(rs app.RequestScope, offset, limit int) ([]models.User, error) {
	users := []models.User{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&users)
	return users, err
}

// Get reads the user with the specified email from the database.
func (dao *UserDAO) FindByEmail(rs app.RequestScope, email string) (*models.User, error) {
	var user models.User
	// if not found, return error
	err := rs.Tx().Select().From("user").Where(dbx.HashExp{"email": email}).One(&user)
	return &user, err
}

// Get reads the user with the specified username from the database.
func (dao *UserDAO) FindByUsername(rs app.RequestScope, username string) (*models.User, error) {
	var user models.User
	// if not found, return error
	err := rs.Tx().Select().From("user").Where(dbx.HashExp{"username": username}).One(&user)
	return &user, err
}

// Get reads the user with the specified column and it's value from the database.
func (dao *UserDAO) FindByField(rs app.RequestScope, column, value string) (*models.User, error) {
	var user models.User
	// if not found, return error
	err := rs.Tx().Select().From("user").Where(dbx.HashExp{column: value}).One(&user)
	return &user, err
}

// Update saves the changes to an user in the database.
func (dao *UserDAO) Save(rs app.RequestScope, user *models.User) error {
	user.BeforeUpdate()
	return rs.Tx().Model(user).Exclude("Id").Update()
}
