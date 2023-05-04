/* Задание 1. Чётные и нечётные
Что нужно сделать
Напишите функцию, которая принимает массив чисел, а возвращает два массива: один из чётных чисел, второй из нечётных.*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNums, oddNums := splitEvenOdd(nums)
	fmt.Println("Even numbers:", evenNums)
	fmt.Println("Odd numbers:", oddNums)
}

func splitEvenOdd(nums []int) ([]int, []int) {
	evenNums := []int{}
	oddNums := []int{}
	for _, num := range nums {
		if num%2 == 0 {
			evenNums = append(evenNums, num)
		} else {
			oddNums = append(oddNums, num)
		}
	}
	return evenNums, oddNums
}
