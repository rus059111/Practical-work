package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	input := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software."

	// Разделяем строку на слова
	words := strings.Fields(input)

	// Считаем количество слов, начинающихся с большой буквы
	count := 0
	for _, word := range words {
		if unicode.IsUpper([]rune(word)[0]) {
			count++
		}
	}

	// Выводим результат
	fmt.Printf("Количество слов, начинающихся с большой буквы: %d", count)
}
