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
	for i := 0; i < len(words); i++ {
		print(<-ch + " ")
	}
}
