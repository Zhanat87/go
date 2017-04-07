package main

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

// go run console/generatePasswordHash/main.go
func main() {
	password := []byte("pass")

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))
}
