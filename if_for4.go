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
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	x = sum
	for {
		if sum%3 != 0 {
			fmt.Println(x)
			break
		}
		min := numbers[0]
		for v := range numbers {
			if v < min {
				min = v
			}
		}
		sum -= min
	}
	if x > 0 {
		fmt.Println(x)
	}
	if x <= 0 {
		fmt.Println("NO")
	}
}
