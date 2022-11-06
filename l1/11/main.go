package main

import "fmt"

func intersectionSet(set1, set2 []int) []int {
	set3 := make([]int, 0)
	um := make(map[int]int)

	// Добавляем элементы первого множества в мап
	for _, el := range set1 {
		if um[el] == 0 {
			um[el] = 1
		}
	}

	// Отмечаем, что элемент есть также(!) во втором множестве
	for _, el := range set2 {
		if um[el] == 1 {
			um[el] = 2
		}
	}

	// Добавляем уникальные элементы в множество пересечений
	for k := range um {
		if um[k] == 2 {
			set3 = append(set3, k)
		}
	}

	return set3
}

func main() {
	set1 := []int{1, 2, 3, 4, 4, 5, 5, 5, 6}
	set2 := []int{4, 5, 5, 6, 7, 7, 8}
	set3 := intersectionSet(set1, set2)
	fmt.Println(set1)
	fmt.Println(set2)
	fmt.Println(set3)
}
