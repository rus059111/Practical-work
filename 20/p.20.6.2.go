package main

import "fmt"

func multiplyMatrices(a [][]int, b [][]int) [][]int {
	m := len(a)
	n := len(b[0])
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < 5; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

func main() {
	a := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}
	b := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
	}
	result := multiplyMatrices(a, b)
	fmt.Println(result)
}
