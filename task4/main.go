// Написать метод, который в консоль выводит таблицу умножения. На вход метод получает число, до которого выводит таблицу умножения.
// В консоли должна появиться таблица. Например, если на вход пришло число 5, то получим:

package main

import (
	"fmt"
	"time"
)

func printMultiplicationTable(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%d\t", i*j)
		}
		fmt.Println()
	}
}

func main() {
	var a int
	fmt.Print("Enter a number: ")
	fmt.Scan(&a)
	n := &a
	time.Sleep(3 * time.Second)
	printMultiplicationTable(*n)
}
