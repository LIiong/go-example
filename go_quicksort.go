package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} 
	temp := arr[0]
	less := make([]int, 0)
	more := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] <= temp {
			less = append(less, arr[i])
		} else{
			more = append(more, arr[i])
		}
	}
	return append(append(quickSort(less),temp), quickSort(more)...)
}
func quickSortTest()  {
	arr := []int{1,4,6,2,9,2,4,6,1}
	arr = quickSort(arr)
	fmt.Println(arr)
}