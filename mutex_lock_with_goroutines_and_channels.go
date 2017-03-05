package main

import (
	"runtime"
	"fmt"
	"os"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	logFile := "./log.txt"
	f, _ := os.Create(logFile)
	f.Close()
	logCh := make(chan string, 50)
	go func() {
		for {
			msg, ok := <- logCh
			if ok {
				f, _ := os.OpenFile(logFile, os.O_APPEND, os.ModeAppend)
				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + " - " + msg)
				f.Close()
			} else {
				break
			}
		}
	}()
	mutex := make(chan bool, 1)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			mutex <- true
			go func() {
				msg := fmt.Sprintf("%d + %d = %d\r\n", i, j, i + j)
				logCh <- msg
				fmt.Print(msg)
				<- mutex
			}()
		}
	}
	fmt.Scanln()
}
