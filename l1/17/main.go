package main

import "fmt"

func BinarySearch(a []int, x int) int {
	left := 0
	right := len(a)

	for left < right {
		middle := (left + right) / 2
		if x == a[middle] {
			return middle
		} else if x > a[middle] {
			left = middle + 1
		} else {
			right = middle
		}
	}

	return -1
}

func main() {
	a := []int{1, 3, 7, 9, 19, 25, 35, 43, 55}
	fmt.Println(BinarySearch(a, 55))
	fmt.Println(BinarySearch(a, 0))
}
