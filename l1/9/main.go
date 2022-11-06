package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	numbers := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	// Отправляем числа в первый канал
	go func(ch1 chan<- int) {
		for i := range numbers {
			ch1 <- numbers[i]
		}

		// Закрываем первый канал когда все числа в него отправлены
		close(ch1)
	}(ch1)

	// Читаем числа из первого канала и отпрвляем квадраты во второй канал
	go func(ch2 chan<- int, ch1 <-chan int) {
		for x := range ch1 {
			ch2 <- x * x
		}

		// Закрываем второй канал когда все числа в него отправлены (когда первый канал закроется)
		close(ch2)
	}(ch2, ch1)

	// Читаем числа из второго канала пока он открыт
	for res := range ch2 {
		fmt.Println(res)
	}
}
