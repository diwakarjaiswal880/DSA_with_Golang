package main

import (
	"fmt"
	
)

func main() {
	a:=[]int{15,6,2,632,64,9,413,64}
	v:=63
	
	fmt.Printf("In given array %v is present:%t",v,linearSearchUnsorted(a,v))
}

func linearSearchUnsorted(data []int, value int) bool {
size := len(data)
for i := 0; i < size; i++ {
if value == data[i] {
return true
}
}
return false
}
