package main

import (
	"fmt"
	"sort"
)

func main() {
	family := []struct {
		Name string
		Age  int
	}{
		{"Alice", 23},
		{"David", 2},
		{"Eve", 2},
		{"Bob", 25},
	}

	// Сортировка по возрасту,
	// сохраняя оригинальный порядок или равные элементы.
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age < family[j].Age
	})
	fmt.Println(family) // [{David 2} {Eve 2} {Alice 23} {Bob 25}]
}
