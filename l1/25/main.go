package main

import (
	"fmt"
	"time"
)

func MySleep(dur time.Duration) {
	start := time.Now()
	for {
		if time.Now().After(start.Add(dur)) {
			return
		}
	}
}

func main() {
	fmt.Println(1)
	MySleep(time.Second * 5)
	fmt.Println(1)
}
