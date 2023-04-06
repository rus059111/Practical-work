package main

import (
	"fmt"
	"math"
)

func main() {
	var x int16 = 10
	var y uint8 = 5
	var z float32 = 2.5

	S := float32(2*x) + float32(math.Pow(float64(y), 2)) - 3/z

	fmt.Println(S)
}
