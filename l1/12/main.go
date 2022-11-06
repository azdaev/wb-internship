package main

import "fmt"

// Функция для проверки наличия строки в срезе строк
func isIn(word string, words []string) bool {
	for _, el := range words {
		if word == el {
			return true
		}
	}
	return false
}

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}
	res := make([]string, 0)
	for _, el := range s {
		// Если слова нет в собственном множестве, добавляем его туда
		if !isIn(el, res) {
			res = append(res, el)
		}
	}
	fmt.Println(res)
}
