package models

type Identity interface {
	GetID()    string
	GetName()  string
	GetEmail() string
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u User) GetID() string {
	return u.ID
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetEmail() string {
	return u.Email
}
