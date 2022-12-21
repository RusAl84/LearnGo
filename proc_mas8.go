package main

import (
	"fmt"
)

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
	min := numbers[0]
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	for _, number := range numbers {
		if number == min {
			fmt.Printf("%d ", max)
		} else if number == max {
			fmt.Printf("%d ", min)
		} else {
			fmt.Printf("%d ", number)
		}
	}
}
