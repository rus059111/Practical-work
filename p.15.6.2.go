package main

import "fmt"

// Реверсируем массив
func reverseArray(arr []int) []int {
	reversed := make([]int, len(arr))
	for i, j := 0, len(arr)-1; i <= j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = arr[j], arr[i]
	}
	return reversed
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Реверсируем массив
	reversedNumbers := reverseArray(numbers)

	// Выводим результат
	fmt.Println(reversedNumbers)
}
