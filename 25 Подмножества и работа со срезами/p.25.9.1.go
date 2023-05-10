/* Цель задания
Написать программу для нахождения подстроки в кириллической подстроке. Программа должна запускаться с помощью команды:

 go run main.go --str "строка для поиска" --substr "поиска"

Для реализации такой работы с флагами воспользуйтесь пакетом flags, а для поиска подстроки в строке вам понадобятся руны.



Что нужно сделать
Спроектировать алгоритм поиска подстроки.
Определить строку и подстроку, используя флаги.
Написать алгоритм реализацию для работы со строками UTF-8 (для этого необходимо воспользоваться рунами). */

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
