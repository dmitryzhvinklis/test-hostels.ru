// Разработайте функцию, которая принимает целое число в качестве аргумента и возвращает строку,
// содержащую это число и слово "компьютер" в нужном склонении по падежам в зависимости от числа.
// Например, при вводе числа 25 функция должна возвращать "25 компьютеров", для числа 41 — "41 компьютер",
// а для числа 1048 — "1048 компьютеров".

package main

import (
	"fmt"
	"strconv"
)

func formatComputerCount(number int) string {
	computerStr := "компьютер"
	if number%10 == 1 && number%100 != 11 {
		computerStr += ""
	} else if number%10 >= 2 && number%10 <= 4 && (number%100 < 10 || number%100 >= 20) {
		computerStr += "а"
	} else {
		computerStr += "ов"
	}

	return strconv.Itoa(number) + " " + computerStr
}

func main() {
	fmt.Println(formatComputerCount(25))   // Output: 25 компьютеров
	fmt.Println(formatComputerCount(41))   // Output: 41 компьютер
	fmt.Println(formatComputerCount(1048)) // Output: 1048 компьютеров
}
