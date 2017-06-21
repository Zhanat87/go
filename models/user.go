package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
	"github.com/Zhanat87/go/helpers"
	"os"
	"github.com/go-ozzo/ozzo-validation/is"
	"database/sql"
)

/*
@link https://marcesher.com/2014/10/13/go-working-effectively-with-database-nulls/
@link http://stackoverflow.com/questions/21151765/json-cannot-unmarshal-string-into-go-value-of-type-int64
@link https://github.com/go-ozzo/ozzo-dbx#null-handling
 */
type User struct {
	Id                 int    `json:"id" db:"id"`
	Username           string `json:"username" db:"username"`
	Email              string `json:"email" db:"email"`
	Password           string `json:"password,omitempty" db:"-"`
	PasswordHash       string `json:"-" db:"password_hash"`
	PasswordResetToken JsonNullString `json:"-" db:"password_reset_token"`
	Avatar             JsonNullString `json:"avatar" db:"avatar"`
	AvatarString       JsonNullString `json:"avatar_string,omitempty" db:"avatar_string"`
	FullName           string `json:"full_name" db:"full_name"`
	Phones             JsonNullString `json:"phones,omitempty" db:"phones"`
	Status             int    `json:"status,string" db:"status"`
	Created_at         string `json:"created_at,omitempty" db:"created_at"`
	Updated_at         string `json:"updated_at,omitempty" db:"updated_at"`
	Provider           JsonNullString `json:"provider,omitempty" db:"provider"`
	ProviderId         JsonNullString `json:"provider_id,omitempty" db:"provider_id"`
}

type UserIdentity struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Avatar       string `json:"avatar,omitempty"`
	AvatarString string `json:"avatar_string,omitempty"`
}

func (m User) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Username, validation.Required, validation.Length(0, 100)),
		validation.Field(&m.Email, validation.Required, validation.Length(0, 100), is.Email),
		validation.Field(&m.Password, validation.Required, validation.Length(4, 100)),
		validation.Field(&m.Status, validation.Required),
	)
}

func (m User) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Username, validation.Required, validation.Length(0, 100)),
		validation.Field(&m.Email, validation.Required, validation.Length(0, 100), is.Email),
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
	return m.Avatar.String
}

func (m User) GetAvatarString() string {
	return m.AvatarString.String
}

func (m *User) BeforeInsert() {
	m.Created_at = time.Now().Format("2006-01-02 15:04:05")
	m.Updated_at = m.Created_at

	hash, _ := m.Hash(m.Password)
	m.PasswordHash = hash

	m.SaveAvatar()
}

func (m *User) BeforeUpdate() {
	m.Updated_at = time.Now().Format("2006-01-02 15:04:05")

	if len(m.Password) > 0 {
		hash, _ := m.Hash(m.Password)
		m.PasswordHash = hash
	}

	m.SaveAvatar()
}

func (m *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(m.PasswordHash), []byte(password))
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
	avatarNotEmpty := m.AvatarString.Valid
	avatarString := m.AvatarString.String
	if m.Avatar.Valid {
		avatarUUID := uuid.NewV4()
		img, err := helpers.SaveImageToDisk("static/users/avatars/", avatarUUID.String(), m.Avatar.String)
		if err != nil {
			panic(err)
		}
		m.AvatarString = JsonNullString{sql.NullString{String:img, Valid:true}}
	} else {
		if avatarNotEmpty == false {
			m.AvatarString = JsonNullString{sql.NullString{String:"", Valid:false}}
		}
		if avatarNotEmpty && (len(avatarString) > 0 && avatarString[0:4] != "http") {
			err := os.Remove("static/users/avatars/" + avatarString)
			if err != nil {
				panic(err)
			}
		}
	}
}
