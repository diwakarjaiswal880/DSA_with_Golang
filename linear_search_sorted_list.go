package main

import (
	"fmt"
)

func main() {
	a := []int{1, 26, 30, 69, 70, 76, 89, 123, 189}
	v := 69

	fmt.Printf("In given array %v is present:%t", v, linearSearchSorted(a, v))
}

func linearSearchSorted(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		} else if value < data[i] {
			return false
		}
	}
	return false
}
