package main

import (
	"fmt"
)

func main() {
	// Колебания температуры
	t := []float64{-30.1, -30, -29.9, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 5, 0.1, 9.9, -0.1, -9.9}

	// Map с результатаами
	tm := make(map[int][]float64)

	for _, val := range t {
		valInt := int(val)
		tm[valInt/10] = append(tm[valInt/10], val)
	}

	fmt.Println(tm)
}
