package main

import (
	"fmt"
)

func main() {
	numbers := map[int]int{}
	for {
		var x int
		if n, _ := fmt.Scan(&x); n == 0 {
			break
		}
		_, ok := numbers[x]
		if !ok {
			numbers[x] = 1
			fmt.Println(x)
		}
	}

}
