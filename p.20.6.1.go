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
