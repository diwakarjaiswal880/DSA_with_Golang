package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 59, 26, 99, 30, 69, 198, 70, 645, 76, 198, 123}
	printRepeating2(a)
}

func printRepeating2(data []int) {
	size := len(data)
	sort.Ints(data) // Sort(data,size)
	fmt.Print("Repeating elements are : ")
	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			fmt.Print(" ", data[i])
		}
	}
	fmt.Println()
}
