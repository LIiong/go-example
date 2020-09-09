package main
import (
	"fmt"
)
/*
* 选择排序
*/
func selectionSort(arr []int){
	for i := 0; i < len(arr); i++ {
		temp := i
		for j := i + 1; j < len(arr); j++ {
			if arr[temp] < arr[j] {
				temp = j
			}		
		}
		arr[i],arr[temp] = arr[temp],arr[i]
	}
}
/*
* 插入排序
*/
func insertionSort(arr []int){
	n := len(arr)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		value := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if value < arr[j] {
				arr[j + 1] = arr[j]
			} else {
				break
			}
		}
		arr[j + 1] = value
	}
}

func selectionTest()  {
	arr := []int{2,3,1,8,3,5,7,2,9}
	insertionSort(arr)
	fmt.Println(arr)
}