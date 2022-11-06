package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1<<55 + 1<<32 + 1<<27)
	b := big.NewInt(1 << 26)
	product := new(big.Int).Mul(a, b)
	division := new(big.Float).Quo(new(big.Float).SetInt(a), new(big.Float).SetInt(b))
	sum := new(big.Int).Add(a, b)
	sub := new(big.Int).Sub(a, b)

	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", b)
	fmt.Printf("a * b = %d\n", product)
	fmt.Printf("a / b = %v\n", division)
	fmt.Printf("a + b = %d\n", sum)
	fmt.Printf("a - b = %d\n", sub)
}
