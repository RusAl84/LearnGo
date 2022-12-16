package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{}
	positions := map[string]int{}
	pos := 0
	for {
		var x string
		if n, _ := fmt.Scan(&x); n == 0 {
			break
		}
		pos += 1
		words = append(words, x)
		positions[x] = pos
	}
	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})
	for _, word := range words {
		fmt.Println(positions[word])
	}
}
