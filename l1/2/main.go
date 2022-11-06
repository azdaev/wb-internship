package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	for i := range nums {
		wg.Add(1)
		go func(x int) {
			defer wg.Done() // Обозначение выполнения горутины
			fmt.Println(x * x)
		}(nums[i])
	}
	wg.Wait() // Программа ждет выполнения всех горутин
}
