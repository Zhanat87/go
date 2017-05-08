package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
	"github.com/Zhanat87/go/helpers"
	"os"
)

/*
@link https://marcesher.com/2014/10/13/go-working-effectively-with-database-nulls/
@link http://stackoverflow.com/questions/21151765/json-cannot-unmarshal-string-into-go-value-of-type-int64
@link https://github.com/go-ozzo/ozzo-dbx#null-handling
 */
type User struct {
	Id           int    `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	Email        string `json:"email" db:"email"`
	Password     string `json:"password,omitempty" db:"password"`
	Avatar       string `json:"avatar" db:"avatar"`
	AvatarString string `json:"avatar_string,omitempty" db:"avatar_string"`
	Full_name    string `json:"full_name" db:"full_name"`
	Phones       *string `json:"phones,omitempty" db:"phones"`
	Status       int8   `json:"status,string" db:"status"`
	Created_at   string `json:"created_at,omitempty" db:"created_at"`
	Updated_at   string `json:"updated_at,omitempty" db:"updated_at"`
}

func (m User) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Username, validation.Required, validation.Length(0, 100)),
		validation.Field(&m.Email, validation.Required, validation.Length(0, 100)),
		validation.Field(&m.Password, validation.Required, validation.Length(4, 100)),
		validation.Field(&m.Status, validation.Required),
	)
}

func (m User) GetId() int {
	return m.Id
}

func (m User) GetUsername() string {
	return m.Username
}

func (m User) GetEmail() string {
	return m.Email
}

func (m User) GetAvatar() string {
	return m.Avatar
}

func (m User) GetAvatarString() string {
	return m.AvatarString
}

func (m *User) BeforeInsert() {
	m.Created_at = time.Now().Format("2006-01-02 15:04:05")
	m.Updated_at = m.Created_at

	hash, _ := m.Hash(m.Password)
	m.Password = hash

	m.SaveAvatar()
}

func (m *User) BeforeUpdate() {
	m.Updated_at = time.Now().Format("2006-01-02 15:04:05")

	if len(m.Password) > 0 {
		hash, _ := m.Hash(m.Password)
		m.Password = hash
	}

	m.SaveAvatar()
}

func (m *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func (m *User) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (m *User) SaveAvatar() {
	avatarString := m.AvatarString
	if len(m.Avatar) > 0 {
		avatarUUID := uuid.NewV4()
		img, err := helpers.SaveImageToDisk("static/users/avatars/", avatarUUID.String(), m.Avatar)
		if err != nil {
			panic(err)
		}
		m.AvatarString = img
	} else {
		m.AvatarString = ""
	}
	if len(avatarString) > 0 {
		err := os.Remove("static/users/avatars/" + avatarString)
		if err != nil {
			panic(err)
		}
	}
}
