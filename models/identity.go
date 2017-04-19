package models

type Identity interface {
	GetId()    int
	GetUsername()  string
	GetEmail() string
}
