package main

import "fmt"

func main() {
	var numbers [4]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 5
	numbers[3] = 3

	//Ввод чисел
	//var x int
	//var result []int
	//for {
	//	n, _ := fmt.Scan(&x)
	//	if n == 0 {
	//		break
	//	}
	//	if x == 0 {
	//		break
	//	}
	//	numbers = append(numbers, x)
	//}

	//fmt.Println("Введеный массив")
	//for i, number := range numbers {
	//	fmt.Printf("i=%d number=%d \n", i, number)
	//}
	length_numbers := len(numbers)
	for i, number := range numbers {
		//fmt.Printf("%d справа числа: ", number)
		x := length_numbers
		for j := i + 1; j < length_numbers; j++ {
			if numbers[j] > number {
				x = numbers[j]
				break
			}
		}
		fmt.Printf("%d ", x)
	}
}
