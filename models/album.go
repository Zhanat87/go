package models

import (
	"github.com/go-ozzo/ozzo-validation"
	//"time"
	//"github.com/Zhanat87/go/db"
)

/*
https://github.com/go-ozzo/ozzo-dbx/issues/23
dbx is not orm and can do relational queries, etc with relations
need select by separate queries,
find parent model and then related models by foreign id(s)
 */
type Album struct {
	Id       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	ArtistId uint   `json:"artistId" db:"artist_id"`
}

// Validate validates the Artist fields.
func (m Album) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required, validation.Length(0, 120)),
		validation.Field(&m.ArtistId, validation.Required),
	)
}

//func (u *Album) BeforeInsert(){
//	u.CreatedAt = time.Now()
//	u.UpdatedAt= u.CreatedAt
//}
//
//func (u *User) BeforeUpdate(){
//	u.UpdatedAt = time.Now()
//}
//
//func (u *User) Create() error{
//	u.BeforeInsert()
//	err:= db.Model(u).Insert()
//	if err != nil{
//		println("Exec err:", err.Error())
//	}
//	return err
//}
