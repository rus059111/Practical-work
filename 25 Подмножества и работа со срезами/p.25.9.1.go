package main

import (
	"flag"
	"fmt"
)

func main() {
	strPtr := flag.String("str", "", "строка для поиска")
	substrPtr := flag.String("substr", "", "подстрока для поиска")
	flag.Parse()

	if *strPtr == "" || *substrPtr == "" {
		fmt.Println("Необходимо указать оба параметра --str и --substr")
		return
	}

	str := []rune(*strPtr)
	substr := []rune(*substrPtr)

	found := false

	for i := 0; i <= len(str)-len(substr); i++ {
		if str[i:i+len(substr)] == substr {
			fmt.Printf("Подстрока найдена в позиции %d\n", i)
			found = true
		}
	}

	if !found {
		fmt.Println("Подстрока не найдена")
	}
}
