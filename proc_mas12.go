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
	length_numbers := len(numbers)
	for i, number := range numbers {
		x := length_numbers
		for j := i + 1; j < length_numbers; j++ {
			//for j := length_numbers - 1; j > i; j-- {
			if numbers[j] > number {
				x = j
				break
			}
		}
		fmt.Println(x)
	}
}
