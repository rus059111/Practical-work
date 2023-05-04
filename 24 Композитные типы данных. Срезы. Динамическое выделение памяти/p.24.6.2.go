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
	result := make([][]int, len(sentences))
	for i, sentence := range sentences {
		words := strings.Split(sentence, " ")
		lastWord := words[len(words)-1]
		lastWordRunes := []rune(lastWord)
		row := make([]int, len(chars))
		for j, char := range chars {
			for k, r := range lastWordRunes {
				if r == char {
					row[j] = k
					break
				}
			}
		}
		result[i] = row
	}
	return result
}

func main() {
	sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := []rune{'H', 'E', 'L', 'П', 'М'}
	result := parseTest(sentences, chars)
	for j, char := range chars {
		fmt.Printf("%c position %d\n", char, result[0][j])
	}
}
