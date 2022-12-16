package main

import "fmt"

func main() {
	var digits string
	fmt.Scan(&digits)
	var revDigits string
	for _, d := range digits {
		revDigits = string(d) + revDigits
	}
	onlyZeros := true
	for _, d := range revDigits {
		if d != '0' {
			onlyZeros = false
		}
		if !onlyZeros {
			fmt.Print(string(d))
		}
	}
	if onlyZeros {
		fmt.Print(0)
	}
	fmt.Println()
}
