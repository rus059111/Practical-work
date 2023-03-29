package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	num1 := 10
	num2 := 20
	fmt.Println("До вызова функции\n", num1, num2)
	swap(&num1, &num2)
	fmt.Println("После вызова функции\n", num1, num2)
}
