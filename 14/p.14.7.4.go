package main

import "fmt"

var global1 = 5
var global2 = 10
var global3 = 15

func addGlobal1(n int) int {
	return n + global1
}

func addGlobal2(n int) int {
	return n + global2
}

func addGlobal3(n int) int {
	return n + global3
}

func main() {
	var n int

	fmt.Print("Введите число: ")
	fmt.Scanln(&n)

	result1 := addGlobal1(n)
	result2 := addGlobal2(result1)
	result3 := addGlobal3(result2)

	fmt.Printf("Результат первой функции: %d\n", result1)
	fmt.Printf("Результат второй функции: %d\n", result2)
	fmt.Printf("Результат третьей функции: %d\n", result3)
}
