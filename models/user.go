package models

import (
	"github.com/go-ozzo/ozzo-validation"
)

type User struct {
	Id         int    `json:"id" db:"id"`
	Username   string `json:"username" db:"username"`
	Email      string `json:"email" db:"email"`
	Password   string `json:"password" db:"password"`
	Avatar     string `json:"avatar" db:"avatar"`
	Full_name  string `json:"full_name" db:"full_name"`
	Phones     string `json:"phones" db:"phones"`
	Status     int8   `json:"status" db:"status"`
	Created_at string `json:"created_at" db:"created_at"`
	Updated_at string `json:"updated_at" db:"updated_at"`
	Deleted_at string `json:"deleted_at" db:"deleted_at"`
}

func (m User) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Username, validation.Required, validation.Length(0, 100)),
		validation.Field(&m.Email, validation.Required, validation.Length(0, 100)),
		validation.Field(&m.Password, validation.Required, validation.Length(6, 100)),
		validation.Field(&m.Status, validation.Required),
	)
}

func (m User) GetId() string {
	return m.Id
}

func (m User) GetUsername() string {
	return m.Username
}

func (m User) GetEmail() string {
	return m.Email
}
