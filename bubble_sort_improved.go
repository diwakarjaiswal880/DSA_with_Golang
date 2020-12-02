package main

import (
	"fmt"
)

func BubbleSort2(arr []int) {
	size := len(arr)
	swapped := 1
	for i := 0; i < (size-1) && swapped == 1; i++ {
		swapped = 0
		for j := 0; j < size-i-1; j++ {
			if arr[j]> arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swapped = 1
			}
		}
	}
}
func main() {
	data := []int{9, 1, 8, 2, -7, 3, 16, 3, 4, 5}
	fmt.Println("Unsorted List : ",data)
	BubbleSort2(data)
	fmt.Println("Sorted List :  ",data)
}
