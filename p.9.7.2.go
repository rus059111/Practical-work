package main

//Задание 2. Минимальный тип данных

import (
	"fmt"
	"math"
)

func main() {
	var Number1, Number2 int16
	var Number3 int64

	var NumberType string

	fmt.Println("Введите первое число:")
	fmt.Scan(&Number1)
	fmt.Println("Введите второе число:")
	fmt.Scan(&Number2)
	Number3 = int64(Number1) * int64(Number2)
	fmt.Println(Number3)

	switch {

	case Number3 <= math.MaxUint8 && Number3 >= 0:
		NumberType = "Uint8"

	case Number3 <= math.MaxUint16 && Number3 >= 0:
		NumberType = "Uint16"
	case Number3 <= math.MaxUint32 && Number3 >= 0:
		NumberType = "Uint32"

	case Number3 >= math.MinInt8 && Number3 <= math.MaxInt8:
		NumberType = "Int8"

	case Number3 >= math.MinInt16 && Number3 <= math.MaxInt16:
		NumberType = "Int16"
	case Number3 >= math.MinInt32 && Number3 <= math.MaxInt32:
		NumberType = "Int32"

	}
	fmt.Println(NumberType)
}
