package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 5
	res := append(a[:i], a[i+1:]...)
	fmt.Println(res)
}
