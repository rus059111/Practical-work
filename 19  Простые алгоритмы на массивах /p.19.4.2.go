package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ { // Проходим по массиву n-1 раз
		for j := 0; j < n-i-1; j++ { // В каждом проходе сравниваем пары соседних элементов
			if arr[j] > arr[j+1] { // Если предыдущий элемент больше следующего, меняем их местами
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{6, 5, 3, 1, 8, 7} // Исходный массив
	fmt.Println("Before sorting:", arr)

	bubbleSort(arr) // Сортируем массив

	fmt.Println("After sorting:", arr)
}
