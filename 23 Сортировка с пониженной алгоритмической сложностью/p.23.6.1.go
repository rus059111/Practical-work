/* Задание 1. Сортировка вставками
Что нужно сделать
Напишите функцию, сортирующую массив длины 10 вставками.

Советы и рекомендации
Алгоритм сортировки доступен на «Википедии». */

package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	arr := []int{5, 2, 7, 1, 9, 6, 8, 3, 10, 4}
	fmt.Println("До сортировки:", arr)
	insertionSort(arr)
	fmt.Println("После сортировки:", arr)
}
