/*
go run cli/syncWaitGroup/main.go
https://golang.org/pkg/sync/
https://stackoverflow.com/questions/19208725/example-for-sync-waitgroup-correct
https://play.golang.org/p/ecvYHiie0P
http://nanxiao.me/en/use-sync-waitgroup-in-golang/
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func dosomething(millisecs time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	duration := millisecs * time.Millisecond
	time.Sleep(duration)
	fmt.Println("Function in background, duration:", duration)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	go dosomething(200, &wg)
	go dosomething(400, &wg)
	go dosomething(150, &wg)
	go dosomething(600, &wg)

	wg.Wait()
	fmt.Println("Done")
}