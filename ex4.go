package main

import "fmt"

func main() {
	var x int
	var massiv []int
	var result []int
	for {
		n, _ := fmt.Scan(&x)
		if n == 0 {
			break
		}
		massiv = append(massiv, x)
	}
	for i := 0; i < len(result); i++ {
		_ = 0
		for j := 0; j < len(result); j++ {
			if massiv[i] == result[j] {
				_ = 1
			}

		}
	}
}
