package main

import (
	"fmt"
)
/**
* 快排
*/
func quickSort1(arr []int) []int {
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
	return append(append(quickSort1(less),temp), quickSort1(more)...)
}
/**
* 快排 原地排序节约空间
*/
func quickSortNew(arr []int, start int, end int){
	if start >= end {
		return
	}
	p := partition1(arr, start, end)
	quickSortNew(arr, start, p - 1)
	quickSortNew(arr, p + 1, end)
}

func partition1(arr []int, start int, end int) int {
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
func quickSort(arr []int, start int, end int)  {
	if start >= end {
		return
	}
	p := partition(arr, start, end)
	quickSort(arr, start, p - 1)
	quickSort(arr, p + 1, end)
}

func partition(arr []int, start int, end int) int {
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
/*
*/
// func removeDuplicates(nums []int) int {

// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	index := 1;
//     for i := 1; i < len(nums); i++ {
// 		flag := false
// 		for j := 0; j < index; j++ {
// 			if nums[j] == nums[i] {
// 				flag = true
// 				break
// 			}
// 		}
// 		if !flag {
// 			nums[index] = nums[i]
// 			index++
// 		}
// 	}
// 	nums = append(nums[:index])
// 	fmt.Println(nums)
// 	return index
// }

func removeDuplicates(nums []int) int {
	for i:=len(nums)-1; i>0; i-- {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	fmt.Println(nums)
	return len(nums)
}

func quickSortTest()  {
	arr := []int{0,0,1,1,1,2,2,3,3,4}
	// arr = quickSort(arr)
	// quickSortNew(arr, 0, len(arr) - 1)
	size := removeDuplicates(arr)
	fmt.Println(arr)
	fmt.Println(size)
}