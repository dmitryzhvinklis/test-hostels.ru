// Написать функцию/метод, которая на вход получает массив положительных целых чисел произвольной длины.
// Например [42, 12, 18]. На выходе возвращает массив чисел, которые являются общими делителями для всех указанных чисел.
// В примере это будет [2, 3, 6].

package main

import (
	"fmt"
)

func findCommonDivisors(numbers []int) []int {
	if len(numbers) == 0 {
		return nil
	}

	// Находим наименьшее число в массиве
	minNumber := numbers[0]
	for _, num := range numbers {
		if num < minNumber {
			minNumber = num
		}
	}

	// Находим общие делители
	commonDivisors := []int{}
	for divisor := 1; divisor <= minNumber; divisor++ {
		isCommonDivisor := true
		for _, num := range numbers {
			if num%divisor != 0 {
				isCommonDivisor = false
				break
			}
		}
		if isCommonDivisor {
			commonDivisors = append(commonDivisors, divisor)
		}
	}

	return commonDivisors
}

func main() {
	numbers := []int{42, 12, 18}
	result := findCommonDivisors(numbers)
	fmt.Println(result)
}
