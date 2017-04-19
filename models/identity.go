package models

type Identity interface {
	GetId()    string
	GetUsername()  string
	GetEmail() string
}
