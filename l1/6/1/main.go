package main

import (
	"fmt"
	"time"
)

// Остановка горутины с помощью канала quit
func main() {
	quit := make(chan bool)
	go func() {
		for {
			select {
			// Если в quit пришло значение, то выполнить данный блок
			case <-quit:
				fmt.Println("quitting")
				return
			default:
				fmt.Println("running")
			}
		}
	}()

	// Позволить горутине выполнять одну секунду
	time.Sleep(1 * time.Second)

	// Отправить значение в канал quit
	quit <- true

	time.Sleep(500 * time.Millisecond)
}
