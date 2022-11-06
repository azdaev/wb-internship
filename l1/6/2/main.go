package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	finished := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				finished <- struct{}{}
			default:
				fmt.Println("Running")
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(2 * time.Second)

		// Отменяем контекст
		cancel()
	}()

	// Блокируем главную горутину чтобы фоновая работала
	<-finished
	fmt.Print("Done")
}
