package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseString(str string) string {

	// Переводим в срез рун чтобы итерироваться по буквам, а не байтам
	sr := []rune(str)
	for i := 0; i < len(sr)/2; i++ {
		sr[i], sr[len(sr)-1-i] = sr[len(sr)-1-i], sr[i]
	}
	return string(sr)
}

func main() {

	// Читаем строку из ввода
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(reverseString(str))
}
