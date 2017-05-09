package helpers

import (
	"fmt"
	"os/user"
	"os"
)

func CurrentUser() *user.User {
	usr, err := user.Current()
	if err != nil {
		SendErrorEmail("golang balu", fmt.Sprintf("failed get current user: %s", err))
		panic(err)
	}
	return usr
}

func IsDocker() bool {
	return os.Getenv("HOME") == "/root"
}
