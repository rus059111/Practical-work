package main

//Задание 2. Дни недели

import (
	"fmt"
)

func main() {

	var word string
	fmt.Println("Введите будний день недели: пн, вт, ср, чт, пт:")
	fmt.Scan(&word)

	switch word {
	case "пн":
		fmt.Println("вторник")
		word = "вт"
		fallthrough
	case "вт":
		fmt.Println("среда")
		word = "чт"
		fallthrough
	case "чт":
		fmt.Println("четвер")
		word = "пт"
		fallthrough
	case "пт":
		fmt.Println("пятница")

	}
}
