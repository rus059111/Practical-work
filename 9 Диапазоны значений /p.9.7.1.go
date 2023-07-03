/* Цель практической работы

Закрепить на практике полученные знания о диапазонах, переполнении и конвертации целочисленных типов.


Что входит в работу

    Переполнение.
    Минимальный тип данных.


Задание 1. Переполнение
Что нужно сделать

В данном модуле мы рассмотрели примеры по целочисленным типам, их размерам в памяти и то, что происходит при её переполнении. Напишите программу, которая в цикле с использованием встроенных констант (на предельные значения целых чисел, в пакете math) будет подсчитывать, сколько приходится переполнений чисел типа uint8, uint16 в диапазоне от 0 до uint32.
Советы и рекомендации

Для нахождения количества переполнений используйте цикл. Воспользуйтесь константами максимальных значений из пакета math.


Задание 2. Минимальный тип данных
Что нужно сделать

Достаточно часто при передаче по Сети или сохранении больших объёмов данных приходится выбирать тип с минимальным размером памяти, чтобы экономить трафик или место на диске. Напишите программу, в которую пользователь вводит два числа (int16), а программа выводит, в какой минимальный тип данных можно сохранить результат умножения этих чисел.
Советы и рекомендации

Обратите внимание, что положительный результат можно сохранить в меньшем типе за счёт использования uint8, uint16. Чтобы не возникло проблем с переполнением в процессе умножения, числа считывайте в int16, а перед умножением переведите их в int32.

Проверить на примерах:

1 1 результат uint8

1 −1 результат int8

640 100 результат uint16

−640 100 результат int32

3000 3000 результат uint32

−3000 3000 результат int32


Что оценивается

    Код программы отформатирован.
    Программа выполняется без ошибок.
    Вычисления выполняются в отдельной строке. */

package main

//Задание 1. Переполнение

import (
	"fmt"
	"math"
)

func main() {

	var valUint8 uint8
	var valUint16 uint16
	var counterUint8 int
	var counterUint16 int

	for i := 0; i <= math.MaxUint32; i++ {
		valUint8 = uint8(i)
		valUint16 = uint16(i)

		if valUint16 == math.MaxUint16 {
			counterUint16++

		}
		if valUint8 == math.MaxUint8 {
			counterUint8++

		}

	}
	fmt.Println("переполнений unit16:", counterUint16-1)
	fmt.Println("переполнений unit8:", counterUint8-1)

}
