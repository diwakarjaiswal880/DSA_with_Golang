package main

import (
	"fmt"
)

func InsertionSort(arr []int) {
	size := len(arr)
	var temp, i, j int
	for i = 1; i < size; i++ {
		temp = arr[i]
		for j = i; j > 0 && arr[j-1]> temp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}

func main() {
	data := []int{9, 1, 8, 2, -7, 3, 16, 3, 4, 5}
	fmt.Println("Unsorted List : ", data)
	InsertionSort(data)
	fmt.Println("Sorted List  : ", data)
}
