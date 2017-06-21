package main

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"flag"
)

// go run cli/generatePasswordHash/main.go -password=test
// https://dhdersch.github.io/golang/2016/01/23/golang-when-to-use-string-pointers.html
// https://gobyexample.com/command-line-flags
func main() {
	passwordFlag := flag.String("password", "pass", "a password")

	flag.Parse()

	password := []byte(*passwordFlag)
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(password) + " hash: " + string(hashedPassword))
}
