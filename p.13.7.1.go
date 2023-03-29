package main

import (
	"fmt"
)

func summ(a, b int) {

	if a > b {
		a, b = b, a
	}

	sum := a
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func main() {
	summ(1, 4)
}
