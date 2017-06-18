package main

import (
	"log"
	"time"
	"os"
	"strconv"
)

func main() {
	for i := 0; i < 10; i++ {
		f, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		log.SetOutput(f)
		log.Print("test " + strconv.Itoa(i))
		time.Sleep(time.Minute)
	}
}
