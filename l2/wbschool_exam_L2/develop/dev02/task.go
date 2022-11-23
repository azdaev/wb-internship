package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Repeat(char rune, n int) []rune {
	res := make([]rune, 0)
	for i := 0; i < n; i++ {
		res = append(res, char)
	}
	return res
}

func Parse(s string) (string, error) {
	runes := []rune(s)        // convert to runes for iterating
	result := make([]rune, 0) // result slice
	last := rune(' ')         // last non-digit character
	for i := 0; i < len(runes); i++ {
		if string(runes[i]) == `\` { // escaping backslash
			if i+1 < len(runes) {
				i++
				last = runes[i]
				result = append(result, last)
				continue
			}
		}
		num, err := strconv.Atoi(string(runes[i])) // check if digit and convert to int
		if err != nil {                            // if non-digit
			last = runes[i]               // save last
			result = append(result, last) // add to result slice
		} else {
			if last == ' ' { // if string starts with digit its invalid
				return "", errors.New("invalid string")
			}
			result = append(result, Repeat(last, num-1)...) // append to res last rune num-1 times
		}
	}
	return string(result), nil
}

func main() {
	a := "a4bc2d5e"
	fmt.Println(a)
	fmt.Println(Parse(a))
}
