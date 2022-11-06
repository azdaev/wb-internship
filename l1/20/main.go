package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReverseSentence(sentence string) string {

	// Срез всех слов из строки
	words := strings.Fields(sentence)
	res := make([]string, 0)

	// Итерируемся по словам с конца и добавляем в конец среза-результата
	for i := len(words) - 1; i >= 0; i-- {
		res = append(res, words[i])
	}

	// Соединяем срез в строку и возвращаем
	return strings.Join(res, " ")
}

func main() {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	fmt.Println(ReverseSentence(line))
}
