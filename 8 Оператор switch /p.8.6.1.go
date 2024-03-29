/* Цель практической работы

Закрепить на практике полученные знания об операторах switch и fallthrough.


Что входит в работу

    Времена года.
    Дни недели.
    По желанию. Расчёт сдачи.

Задание по желанию — необязательная, повышенной сложности, требует самостоятельного изучения дополнительного материала.


Задание 1. Времена года
Что нужно сделать

Пользователь вводит месяц, программа должна вывести, на какое время года (зиму, весну, лето, осень) этот месяц выпадает.

Как группировать:

    декабрь, январь, февраль — зима;
    март, апрель, май — весна;
    июнь, июль, август — лето;
    сентябрь, октябрь, ноябрь — осень.

Рекомендация

Пример работы программы:

Времена года.
Введите месяц:
январь
Время года — зима.


Задание 2. Дни недели
Что нужно сделать

Пользователь вводит будний день недели в сокращённой форме (пн, вт, ср, чт, пт) и получает развёрнутый список всех последующих рабочих дней, включая пятницу.
Рекомендация

Пример работы программы:

Дни недели.
Введите будний день недели: пн, вт, ср, чт, пт:
вт
вторник
среда
четверг
пятница


Задание 3 (по желанию). Расчёт сдачи
Что нужно сделать

Напишите функцию, которая посчитает, сможет ли продавец в киоске обслужить всех покупателей. В киоске каждый лимонад стоит пять долларов. Клиенты стоят в очереди, чтобы купить у вас, и заказывают по одному лимонаду. Каждый покупатель может купить только один лимонад и заплатить купюрами номиналом 5, 10 или 20 долларов. Вы должны дать каждому покупателю сдачу с его купюры.

Обратите внимание, что сначала у вас нет сдачи.
Советы и рекомендации

Сигнатура функции lemonadeChange(bills []int) bool,

где bills — это купюры, которые мы получаем от покупателей, по одной купюре от каждого.

Верните true, если вы можете предоставить каждому покупателю правильную сдачу.

Пример 1

Расчёт сдачи.
Ввод: [5,5,5,10,20]
Вывод: true

Пояснение:

От первых трёх клиентов мы собираем три купюры по 5 долларов по порядку.

От четвёртого покупателя мы получаем купюру номиналом 10 долларов и возвращаем 5 долларов сдачи.

От пятого клиента мы получаем 20 долларов и возвращаем купюры номиналом 10 и 5 долларов.

Поскольку все клиенты получили правильную сдачу, мы выводим true.

Пример 2

Расчёт сдачи.
Ввод: [10,10]
Вывод: false

Пояснение:

Поскольку купюр в кассе нет, мы не можем дать сдачу первому покупателю.

Пример 3

Расчёт сдачи.
Ввод: [5,5,10,10,20]
Вывод: false

Пояснение:

От первых двух клиентов по порядку мы получаем две купюры по 5 долларов.

Для следующих двух клиентов мы собираем купюры по 10 долларов и возвращаем 5 долларов третьему и четвёртому клиентам.

Последнему покупателю мы не можем вернуть сдачу в размере 15 долларов, потому что у нас есть только две купюры по 10 долларов.

Поскольку не все покупатели получили правильную сдачу, ответ — false.


Что оценивается

    Код программы отформатирован.
    Программа выполняется без ошибок.
    Вычисления выполняются в отдельной строке. */

package main

//Задание 1. Времена года

import (
	"fmt"
)

func main() {

	var word string
	fmt.Println("Введите месяц:")
	fmt.Scan(&word)

	switch word {
	case "декабрь", "январь", "февраль":
		fmt.Println("зима")

	case "март", "апрель", "май":
		fmt.Println("весна")
	case "июнь", "июль", "авгус":
		fmt.Println("лето")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println("осень")
	default:
		fmt.Println("нет такого месяца")
	}
}
