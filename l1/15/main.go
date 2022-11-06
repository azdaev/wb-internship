package main

import "fmt"

// Надо избавиться от глобальной переменной
// str[:100] может вернуть не 100 символов, а меньше, так как не все символы занимают один байт

// Возможная имплементация функции createHugeString
func createHugeString(size int) string {
	s := make([]rune, size)
	for i := range s {
		s[i] = 'ж'
	}
	return string(s)
}

func someFunc() string {
	v := createHugeString(1 << 10)

	// Переводим в срез рун, так как он безопасен в плане разделения строки на символы, а не на байты
	vr := []rune(v)
	return string(vr[:100])
}
func main() {
	justString := someFunc()
	fmt.Println(justString, len([]rune(justString)))
}
