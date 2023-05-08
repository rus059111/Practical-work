package main

import (
	"flag"
	"fmt"
	//"strings"
)

func main() {
	strPtr := flag.String("strRunes", "Строка для поиска", "a string for searching")
	substrPtr := flag.String("substrRunes", "поиска", "a substring to search for")
	flag.Parse()

	if *strPtr == "" || *substrPtr == "" {
		fmt.Println("Не заполнена строка и подстрока.")

		return
	}

	strRunes := []rune(*strPtr)
	substrRunes := []rune(*substrPtr)

	// ищем индекс первого вхождения подстроки в строку
	index := -1
	for i := range strRunes {
		if i+len(substrRunes) > len(strRunes) {
			break
		}
		if string(strRunes[i:i+len(substrRunes)]) == *substrPtr {
			index = i
			//break
		}
	}

	if index == -1 {
		fmt.Println("Подстрока не найдена")
		return
	}

	fmt.Printf("Подстрока найдена, c index %d \n", index)

}
