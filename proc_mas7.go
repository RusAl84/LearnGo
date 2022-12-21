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
	chet := []int{}
	ne_chet := []int{}
	for _, number := range numbers {
		if number%2 == 0 {
			chet = append(chet, number)
		} else {
			ne_chet = append(ne_chet, number)
		}
	}
	for i := len(ne_chet) - 1; i >= 0; i-- {
		fmt.Printf("%d ", ne_chet[i])
	}
	for _, number := range chet {
		fmt.Printf("%d ", number)
	}
}
