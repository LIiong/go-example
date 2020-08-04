package main
import (
	"fmt"
)
var muns = []int{1,2,3,4,5,6,7,8,9}
/*
* 二分查找
*/
func binarySearch(v int) int{
	start, end := 0, len(muns) - 1
	for start <= end {
		fmt.Printf("start:%d,end:%d\n", start, end)
		mid := (start + end)/2
		if muns[mid] > v {
			end = mid - 1
		} else if muns[mid] < v {
			start = mid + 1
		} else {
			return mid;
		}
	}
	fmt.Printf("start:%d,end:%d\n", start, end)
	return -1
}
/*
* 递归二分查找
*/
func binarySearch2(v int){

	start,end := 0,len(muns) - 1
	s := binarySearchRecursion(start, end, v)
	fmt.Println(s)
}
func binarySearchRecursion(start int, end int, v int) int {
	mid := (start + end)/2
	if muns[mid] > v {
		end = mid - 1
	} else if muns[mid] < v {
		start = mid + 1
	} else {
		return mid;
	}
	if start <= end {
		return binarySearchRecursion(start, end, v)
	}
	return -1
}