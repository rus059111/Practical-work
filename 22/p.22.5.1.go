package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Инициализируем генератор случайных чисел

	arr := make([]int, 10) // Создаем массив из 10 элементов

	// Заполняем массив случайными числами от 0 до 99
	for i := range arr {
		arr[i] = rand.Intn(100)
	}

	fmt.Println("Массив:", arr) // Выводим массив на экран

	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num) // Считываем введенное число

	found := false
	count := 0
	for _, v := range arr {
		if v == num {
			found = true
		}
		if found {
			count++
		}
	}

	if found {
		fmt.Printf("Чисел после %d: %d\n", num, count-1)
	} else {
		fmt.Println("Число не найдено")
	}
}
