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

func selectionTest()  {
	arr := []int{2,3,1,8,3,5,7,2,9}
	selectionSort(arr)
	fmt.Println(arr)
}