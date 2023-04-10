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
