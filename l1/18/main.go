package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	val int
	sync.Mutex
}

func (c *Counter) Increment() {
	defer c.Unlock()
	c.Lock()
	c.val++
}

func main() {
	wg := &sync.WaitGroup{}
	c := Counter{
		val: 0,
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}

	wg.Wait()
	fmt.Print(c.val)
}
