package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	search := 6
	index := binary_search(arr,search)
	fmt.Println(index)

	low := 0
	high := len(arr) - 1
	index2 := binary_search_recursoin(arr,low,high,search)
	fmt.Println(index2)

	aaa := make(map[string]string)
}

//二分查找
func binary_search(arr []int, searchValue int) int {
	low := 0
	high := len(arr) -1
	for low <= high {
		mid := low + (high - low) / 2
		if arr[mid] == searchValue {
			return mid
		}else if arr[mid] > searchValue {
			high = mid - 1
		}else if arr[mid] < searchValue {
			low = mid + 1
		}
	}
	return -1
}

//二分查找递归实现
func binary_search_recursoin(arr []int,low,high,searchValue int) int {
	mid := low + (high - low)/2
	if arr[mid] == searchValue {
		return mid
	}else if arr[mid] > searchValue {
		high = mid - 1
	}else if arr[mid] < searchValue {
		low = mid + 1
	}
	return binary_search_recursoin(arr,low,high,searchValue)
}