package main

import "fmt"

func main() {
	a, b := 1, 2
	fmt.Printf("a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("a = %d, b = %d", a, b)

}
