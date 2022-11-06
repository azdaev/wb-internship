package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
)

func main() {
	var N int

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		for sig := range signalChannel {
			fmt.Printf("%d workers were killed when captured '%v' signal caused by user\n", N, sig)
			os.Exit(0)
		}
	}()

	_, err := fmt.Scan(&N)
	if err != nil {
		return
	}

	c := make(chan int)

	for i := 0; i < N; i++ {
		go func() {
			for {
				fmt.Println(<-c)
			}
		}()
	}

	for {
		c <- rand.Int()
	}

}
