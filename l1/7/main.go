package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make(map[int]int) // Наша map
	mutex := &sync.RWMutex{}  // RWMutex (нет необходимости в обычном Mutex)
	w := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		w.Add(1)
		go func(x int) {
			defer w.Done()
			mutex.Lock() // Блокируем запись
			data[x] = x * x
			mutex.Unlock() // Разблокируем запись
		}(i)
	}

	w.Wait()
	fmt.Println(data)
}
