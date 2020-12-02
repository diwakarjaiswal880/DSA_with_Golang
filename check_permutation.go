package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 59, 26, 99, 30, 69, 198, 70, 645, 76, 198, 123}
	b := []int{1, 59, 26, 99, 30, 69, 198, 70, 645, 76, 198, 123}
	c := CheckPermutation(a, b)
	fmt.Printf("Given two lists are in permutation: %t", c)
}

func CheckPermutation(data1 []int, data2 []int) bool {
	size1 := len(data1)
	size2 := len(data2)
	if size1 != size2 {
		return false
	}
	sort.Ints(data1)
	sort.Ints(data2)
	for i := 0; i < size1; i++ {
		if data1[i] != data2[i] {
			return false
		}
	}
	return true
}
