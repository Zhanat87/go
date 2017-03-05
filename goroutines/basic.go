package main

import (
	"time"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	go func() {
		for i := 0; i < 100; i++ {
			println("Hello")
			time.Sleep(10 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			println("Go")
			time.Sleep(10 * time.Millisecond)
		}
	}()
	// delay for main routine
	time.Sleep(time.Second)
	println(runtime.NumCPU())
}
