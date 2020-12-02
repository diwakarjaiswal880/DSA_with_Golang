package main

import (
	"fmt"
)

func main() {
	a := []int{1, 59, 26, 99, 30, 69, 198, 70, 645, 76, 198, 123}
	printRepeating(a)
}

func printRepeating(data []int) {
	size := len(data)
	fmt.Print("Repeating elements are : ")
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				fmt.Print(" ", data[i])
			}
		}
	}
	fmt.Println()
}
