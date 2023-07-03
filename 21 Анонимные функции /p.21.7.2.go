package main

import "fmt"

func wrapper(f func(int, int) int, a, b int) int {
	defer func() {
		fmt.Println("Функция выполнена")
	}()
	return f(a, b)
}

func main() {
	result1 := wrapper(func(a, b int) int {
		return a + b
	}, 2, 3)

	result2 := wrapper(func(a, b int) int {
		return a * b
	}, 4, 5)

	result3 := wrapper(func(a, b int) int {
		return a - b
	}, 6, 3)

	fmt.Println(result1, result2, result3)
}
