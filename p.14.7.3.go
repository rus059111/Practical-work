package main

import (
	"fmt"
)

func multiplyNumber(n int, multiplier int) (result int) {
	result = n * multiplier
	return
}

func addNumber(n int, addend int) (result int) {
	result = n + addend
	return
}

func main() {
	var n, multiplier, addend int

	fmt.Print("Введите число: ")
	fmt.Scanln(&n)

	fmt.Print("Введите множитель: ")
	fmt.Scanln(&multiplier)

	fmt.Print("Введите слагаемое: ")
	fmt.Scanln(&addend)

	result1 := multiplyNumber(n, multiplier)
	result2 := addNumber(result1, addend)

	fmt.Printf("Результат первого преобразования: %d\n", result1)
	fmt.Printf("Результат второго преобразования: %d\n", result2)
}
