package main

import (
	"fmt"
)

func sum(numbers []int, c chan<- int) {
	sum := 0
	for _, num := range numbers {
		sum += num * num
	}
	c <- sum // Отправить результат в канал
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	c1 := make(chan int)
	c2 := make(chan int)
	go sum(numbers[:len(numbers)/2], c1) // Первую половину суммирует одна горутина
	go sum(numbers[len(numbers)/2:], c2) // Втору половину суммирует вторая горутина
	fmt.Println(<-c1 + <-c2)
}

//func main() {
//	t1 := time.Now()
//	s := []int{2, 4, 6, 8, 10}
//
//	c := make(chan int)
//	go sum(s, c)
//
//	fmt.Println(<-c)
//	fmt.Println(time.Since(t1))
//}

//func main() {
//	t1 := time.Now()
//	s := []int{2, 4, 6, 8, 10}
//	sum := 0
//	for _, v := range s {
//		sum += v * v
//	}
//	fmt.Println(sum)
//	fmt.Println(time.Since(t1))
//}
