package main

import "fmt"

func main() {
	var numbers [10]int
	var evenCount, oddCount int

	// Запрашиваем у пользователя 10 целых чисел
	for i := 0; i < 10; i++ {
		fmt.Printf("Введите число %d: ", i+1)
		fmt.Scan(&numbers[i])
	}

	// Подсчитываем количество четных и нечетных чисел
	for _, num := range numbers {
		if num%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	// Выводим результат
	fmt.Printf("Четных чисел: %d\n", evenCount)
	fmt.Printf("Нечетных чисел: %d\n", oddCount)
}
