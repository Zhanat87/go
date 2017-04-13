package db

import (
	"sync"
	"github.com/jinzhu/gorm"
)

type singleton gorm.DB

var (
	once sync.Once

	instance singleton
)

/*
https://github.com/tmrts/go-patterns/blob/master/creational/singleton.md
http://stackoverflow.com/questions/41257847/how-to-create-singleton-db-class-in-golang
note: best decision: it's call db open every time
 */
func New() singleton {
	// call only one time
	once.Do(func() {
		instance, err := gorm.Open("postgres", "host=localhost user=postgres dbname=go_restful sslmode=disable password=postgres")
		// can't call defer close here
		defer instance.Close()
		if err != nil {
			return
		}
	})

	return instance
}