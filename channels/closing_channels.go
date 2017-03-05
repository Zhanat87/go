package main

import (
	"strings"
)

func main() {
	phrase := "that a text thet need explode by probel.\r\n"
	words := strings.Split(phrase, " ")
	ch := make(chan string, len(words))
	for _, word := range words {
		ch <- word
	}
	close(ch)
	for msg := range ch {
		print(msg + " ")
	}
	//for {
	//	if msg, ok := <- ch; ok {
	//		print(msg + " ")
	//	} else {
	//		break
	//	}
	//}
}
