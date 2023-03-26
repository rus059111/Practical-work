package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// Разделяем строку на отдельные части
	parts := strings.Fields(input)

	// Ищем все части, которые можно привести к числу в десятичном формате
	for _, part := range parts {
		if num, err := strconv.Atoi(part); err == nil {
			fmt.Println(num)
		} else if num, err := strconv.ParseFloat(part, 64); err == nil {
			fmt.Println(num)
		}
	}
}
