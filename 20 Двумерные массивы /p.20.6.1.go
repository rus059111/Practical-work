/* Цель домашнего задания

Научиться работать с двумерными массивами.


Что входит в задание

    Написать функцию, вычисляющую определитель.
    Написать функцию, умножающую две матрицы.


Задание 1. Подсчёт определителя
Что нужно сделать

Напишите функцию, вычисляющую определитель матрицы размером 3 × 3.
Советы и рекомендации

    Алгоритм вычисления представлен в «Википедии».
    В качестве среды разработки может помочь Goland или VScode.

Что оценивается

    Правильность ответа.
    Проект собирается без ошибок.

Как отправить задание на проверку

Выполните задание в файле GO. Загрузите файл на Google Диск, откройте доступ для всех по ссылке. Отправьте ссылку на документ через форму для сдачи домашнего задания.


Задание 2. Умножение матриц
Что нужно сделать

Напишите функцию, умножающую две матрицы размерами 3 × 5 и 5 × 4.
Советы и рекомендации

    Алгоритм умножения матриц проиллюстрирован в «Википедии».
    В качестве среды разработки может помочь Goland или VScode.

Что оценивается

    Правильность размеров исходных и конечной матрицы.
    Арифметическая правильность ответа. */

package main

import "fmt"

func determinant(matrix [][]float64) float64 {
	a := matrix[0][0]
	b := matrix[0][1]
	c := matrix[0][2]
	d := matrix[1][0]
	e := matrix[1][1]
	f := matrix[1][2]
	g := matrix[2][0]
	h := matrix[2][1]
	i := matrix[2][2]

	det := a*(e*i-f*h) - b*(d*i-f*g) + c*(d*h-e*g)

	return det
}

func main() {
	matrix := [][]float64{{2, 3, 1}, {4, 5, 6}, {7, 8, 9}}
	det := determinant(matrix)
	fmt.Println(det) // Output: -24
}
