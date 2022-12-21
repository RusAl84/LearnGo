package main

import "fmt"

func main() {
	freqs := make([]int, 100001)
	for {
		var x int
		if n, _ := fmt.Scan(&x); n == 0 {
			break
		}
		freqs[x] += 1
	}
	for num, hits := range freqs {
		if hits > 0 {
			fmt.Printf("%d => %d \n", num, hits)
		}
	}
}
