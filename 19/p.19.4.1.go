package main

import (
	"fmt"
)

func merge(arr1 []int, arr2 []int) []int {
	merged := make([]int, 0, len(arr1)+len(arr2)) // Создаем новый массив для объединенных значений, длиной len(arr1) + len(arr2), с capacity равным его длине.
	i, j := 0, 0                                  // Инициализируем указатели на начало каждого массива
	for i < len(arr1) && j < len(arr2) {          // Пока не достигнут конец любого из массивов
		if arr1[i] < arr2[j] { // Добавляем минимальный элемент в объединенный массив
			merged = append(merged, arr1[i])
			i++ // Инкрементируем указатель на элемент в первом массиве
		} else {
			merged = append(merged, arr2[j])
			j++ // Инкрементируем указатель на элемент во втором массиве
		}
	}
	// Добавляем оставшиеся элементы из массива arr1, если они есть
	for i < len(arr1) {
		merged = append(merged, arr1[i])
		i++
	}
	// Добавляем оставшиеся элементы из массива arr2, если они есть
	for j < len(arr2) {
		merged = append(merged, arr2[j])
		j++
	}
	return merged // Возвращаем объединенный массив

}
func main() {

	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{2, 4, 6, 8, 9}

	merged := merge(arr1, arr2)

	fmt.Println(merged) // [1 2 3 4 5 6 7 8 9]

}
