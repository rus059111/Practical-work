package main

//Задание 1. Разложение ex в ряд Тейлора

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	var n int

	// Вводим значение x и точность вычисления
	fmt.Print("Введите значение x: ")
	fmt.Scan(&x)

	fmt.Print("Введите точность вычисления (количество знаков после запятой): ")
	fmt.Scan(&n)

	// Вычисляем значение экспоненты
	sum := 1.0
	term := 1.0
	i := 1

	for {
		term *= x / float64(i)
		sum += term

		if math.Abs(term) < math.Pow10(-n) {
			break
		}

		i++
	}

	// Выводим результат
	fmt.Printf("exp(%.2f) с точностью %d знаков после запятой: %.15f", x, n, sum)
}
