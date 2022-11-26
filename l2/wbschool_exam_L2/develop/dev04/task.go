package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// SortString - функция для сортировки букв в строке
func SortString(word string) string {
	letters := []rune(word)
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})
	return string(letters)
}

func GroupAnagrams(words []string) *map[string][]string {
	// Словарь для избежания дубликатов
	dictionary := make(map[string]bool, len(words))
	// Мапа - предрезультат
	m := make(map[string][]string)
	// Мапа - результат
	r := make(map[string][]string)

	// Итерируемся по словам
	for _, upperWord := range words {

		// Переводим слово в нижний регистр
		word := strings.ToLower(upperWord)
		if len(word) < 2 {
			continue // Пропускаем слова из одной буквы
		}

		// Если слова уже нет в каком-либо множестве
		if !dictionary[word] {
			m[SortString(word)] = append(m[SortString(word)], word)
			dictionary[word] = true
		}
	}

	for _, v := range m {
		firstWord := v[0]
		sort.Strings(v)
		r[firstWord] = v
	}

	return &r
}

func main() {
	fmt.Print(*GroupAnagrams([]string{"eat", "ate", "tea", "bike", "kibe", "cab", "bca", "abc", "abc"}))
}
