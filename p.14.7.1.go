package main

import (
	"fmt"
)

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var num int
	fmt.Println("Введите число:")
	fmt.Scan(&num)

	if isEven(num) {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}
}
