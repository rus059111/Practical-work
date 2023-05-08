/* Задание 2. Поиск символов в нескольких строках
Что нужно сделать
Напишите функцию, которая на вход принимает массив предложений (длинных строк) и массив символов типа rune, а возвращает 2D-массив, где на позиции [i][j] стоит индекс вхождения символа j из chars в последнее слово в предложении i (строку надо разбить на слова и взять последнее). То есть сигнатура следующая:

func parseTest(sentences []string, chars []rune)

Советы и рекомендации
В качестве среды разработки используйте Goland или VScode.
Не забудьте проверить, что вы получили больше чем 0 аргументов.
Подход не важен: можно переписать сортировку пузырьком или отсортировать, а потом перевернуть.
Пример входных данных
sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}

chars := [5]rune{'H','E','L','П','М'}

Пример вывода результата в первом элементе массива

'H' position 0

'E' position 1

'L' position 9 */

package main

import (
	"fmt"
	"strings"
)

func parseTest(sentences []string, chars []rune) [][]int {
	result := [][]int{}
	for _, sentence := range sentences {
		words := strings.Fields(sentence)
		lastWord := words[len(words)-1]
		lastIndex := []int{}
		for _, char := range chars {
			index := strings.IndexRune(lastWord, char)

			lastIndex = append(lastIndex, index)
		}
		result = append(result, lastIndex)
	}
	return result
}

func main() {
	sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := []rune{'H', 'E', 'L', 'П', 'М'}
	result := parseTest(sentences, chars)
	/* 	for i := range result[0] {
		fmt.Printf("%c position %d\n", chars[i], result[0][i])
	} */
	for i, indices := range result {
		fmt.Println("Предложение ", sentences[i], indices, "\n")
	}
}
