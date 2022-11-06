package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 5

func main() {
	c := make(chan int)

	go func(c chan<- int) {
		for {
			c <- int(rand.Intn(50))
		}
	}(c)

	go func(c <-chan int) {
		for v := range c {
			fmt.Print(v, " ")
		}
	}(c)

	time.Sleep(N * time.Second)
}
