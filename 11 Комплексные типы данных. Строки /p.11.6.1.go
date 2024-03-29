/* Цель практической работы

Закрепить на практике полученные знания работы со строками и приведением типа.


Что входит в работу

    Определение количества слов, начинающихся с большой буквы.
    Вывод чисел, содержащихся в строке.


Задание 1. Определение количества слов, начинающихся с большой буквы
Что нужно сделать

Напишите программу, которая выведет количество слов, начинающихся с большой буквы в строке: Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software.
Рекомендация

Пример работы программы:

Определение количества слов, начинающихся с большой буквы в строке:
Go is an Open source programming Language that makes it Easy
to build simple, reliable, and efficient Software
Строка содержит 5 слов с большой буквы.


Задание 2. Вывод чисел, содержащихся в строке
Что нужно сделать

Напишите программу, которая выведет все части строки

a10 10 20b 20 30c30 30 dd,

которые можно привести к числу в десятичном формате.
Рекомендация

Пример работы программы:

Исходная строка:
a10 10 20b 20 30c30 30 dd
В строке содержатся числа в десятичном формате:
10 20 30


Что оценивается

    Код программы отформатирован.
    Программа выполняется без ошибок.
    Вычисления выполняются в отдельной строке. */

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
