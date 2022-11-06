package main

import "fmt"

func setBit(num *int64, index uint8, bit uint8) {
	fmt.Printf("%b\n", *num)
	if bit == 1 {
		*num = *num | (1 << index) // 1 << 5 = 100000
	} else if bit == 0 {
		*num = *num & (1<<63 - 1 - (1 << index)) // 1<<63 - 1 - (1 << 5) = 1...11011111.
	} else {
		fmt.Println("wrong bit value. number not changed")
		return
	}
	fmt.Printf("%b\n", *num)
}

func main() {
	var num int64 = 1<<63 - 1
	setBit(&num, 61, 0)
}
