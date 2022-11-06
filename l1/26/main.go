package main

import (
	"fmt"
	"strings"
)

func isUnique(str string) bool {

	// Будем заполнять мапу в цикле (встречалась ли уже очередная буква в строке)
	m := make(map[string]bool)
	for _, letter := range str {

		// Переводим букву в нижний регистр, чтобы функция была регистронезависимой
		letterLower := strings.ToLower(string(letter))

		if m[letterLower] == true { // Буква уже встречалась
			return false
		} else {
			m[letterLower] = true
		}
	}
	return true
}

func main() {
	a := "aabcd"
	fmt.Println(isUnique(a))
}
