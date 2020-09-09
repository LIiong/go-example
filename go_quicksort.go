package main

import (
	"fmt"
)
/**
* 快排
*/
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
/**
* 快排 原地排序节约空间
*/
func quickSortNew(arr []int, start int, end int){
	if start >= end {
		return
	}
	p := partition(arr, start, end)
	quickSortNew(arr, start, p - 1)
	quickSortNew(arr, p + 1, end)
}

func partition(arr []int, start int, end int) int {
	pivot := arr[end]
	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			arr[i],arr[j] = arr[j],arr[i]
			i++
		}
	}
	arr[i],arr[end] = arr[end], arr[i]
	return i
}
/**
*增加哨兵 减少交换次数
*/
func quickSortNew1(arr []int, start int, end int)  {
	if start >= end {
		return
	}
	p := partitionNew(arr, start, end)
	quickSortNew1(arr, start, p - 1)
	quickSortNew1(arr, p + 1, end)
}

func partitionNew(arr []int, start int, end int) int {
	i := start
	j := end
	pivot := arr[i]
	for i < j {
		for i < j &&  pivot <= arr[j]{
			j--
		}
		arr[i] = arr[j]
		for i < j &&  pivot >= arr[i] {
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = pivot
	return i
}

func quickSortTest()  {
	arr := []int{1,4,6,2,9,2,4,6,1}
	// arr = quickSort(arr)
	// quickSortNew(arr, 0, len(arr) - 1)
	quickSortNew1(arr, 0, len(arr) - 1)
	fmt.Println(arr)
}