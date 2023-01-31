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
