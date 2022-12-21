package main

import "fmt"

func main() {
	var numbers []int
	var x int
	for {
		n, _ := fmt.Scan(&x)
		if n == 0 {
			break
		}
		numbers = append(numbers, x)
	}
	n := len(numbers)
	for i := 0; i < n; i++ {
		if numbers[i] < numbers[len(numbers)-i-1] {
			numbers[i], numbers[len(numbers)-i-1] = numbers[len(numbers)-i-1], numbers[i]
		}
	}
	for _, number := range numbers {
		fmt.Printf("%d ", number)
	}
}
