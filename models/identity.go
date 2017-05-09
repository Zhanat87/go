package models

type Identity interface {
	GetId() int
	GetUsername() string
	GetEmail() string
	GetAvatar() string
	GetAvatarString() string
}
