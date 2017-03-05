package main

import (
	"runtime"
	"fmt"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)
	mutex := new(sync.Mutex)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			mutex.Lock()
			go func() {
				mutex.Unlock()
				fmt.Printf("%d + %d = %d\r\n", i, j, i + j)
			}()
		}
	}
	fmt.Scanln()
}
