/* Задание 2. Нахождение первого вхождения числа в упорядоченном массиве (числа могут повторяться)
Что нужно сделать
Заполните упорядоченный массив из 12 элементов и введите число. Необходимо реализовать поиск первого вхождения заданного числа в массив. Сложность алгоритма должна быть минимальная.

Что оценивается
Верность индекса.

При вводе массива 1 2 2 2 3 4 5 6 7 8 9 10 и вводе числа 2 программа должна вывести индекс 1. */

package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	index := -1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			index = mid
			right = mid - 1 // продолжаем поиск в левой части массива
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return index
}

func main() {
	arr := []int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 2

	index := binarySearch(arr, target)

	if index == -1 {
		fmt.Println("Элемент не найден")
	} else {
		fmt.Printf("Первое вхождение элемента %d находится на индексе %d\n", target, index)
	}
}
