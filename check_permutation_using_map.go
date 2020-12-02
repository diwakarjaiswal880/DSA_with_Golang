package main

import (
	"fmt"
)

func main() {
	a := []int{1, 59, 26, 99, 30, 69, 198, 70, 645, 76, 198, 123}
	b := []int{1, 59, 26, 99, 30, 69, 198, 70, 645, 76, 198, 123}
	c := checkPermutation2(a, b)
	fmt.Printf("Given two lists are in permutation: %t", c)
}

func checkPermutation2(data1 []int, data2 []int) bool {
	size1 := len(data1)
	size2 := len(data2)
	h := make(map[int]int)
	if size1 != size2 {
		return false
	}
	for i := 0; i < size1; i++ {
		h[data1[i]]++
		h[data2[i]]--
	}
	for i := 0; i < size1; i++ {
		if h[data1[i]] != 0 {
			return false
		}
	}
	return true
}
