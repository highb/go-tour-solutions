package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var split = strings.Fields(s)
	var count = make(map[string]int)
	for _, word := range split {
	  count[word]++
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
