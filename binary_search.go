package main

import (
	"fmt"
)

func main() {
	a := []int{1, 59, 26, 99, 30, 69, 198, 70, 645, 76, 89, 123}
	v := 645

	fmt.Printf("In given array %v is present:%t", v, Binarysearch(a, v))
}

func Binarysearch(data []int, value int) bool {
	size := len(data)
	low := 0
	high := size - 1
	mid := 0
	for low <= high {
		mid = low + (high-low)/2
		if data[mid] == value {
			return true
		} else if data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
