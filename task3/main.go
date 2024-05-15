// Написать функцию/метод, которая возвращает массив простых чисел в диапазоне (2 числа - минимальное и максимальное) заданных чисел.
// Например, на вход переданы 2 числа: от 11 до 20  (диапазон считается включая граничные значения).
// На выходе программа должна выдать [11, 13 , 17, 19]

package main

import (
	"fmt"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func findPrimesInRange(min, max int) []int {
	primes := []int{}
	for i := min; i <= max; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	min := 11
	max := 20
	result := findPrimesInRange(min, max)
	fmt.Println(result)
}
