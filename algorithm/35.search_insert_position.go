package main

import "fmt"

/**
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
你可以假设数组中无重复元素。
示例 1:
输入: [1,3,5,6], 5
输出: 2

示例 2:
输入: [1,3,5,6], 2
输出: 1

示例 3:
输入: [1,3,5,6], 7
输出: 4

示例 4:
输入: [1,3,5,6], 0
输出: 0
 */
func main() {
	numbers := []int{1,2,3,5,6}
	target := 9
	fmt.Println(searchInsert(numbers,target))
	fmt.Println(searchInsert2(numbers,target))
}
//遍历匹配
func searchInsert(nums []int, target int) int {
	i := 0
	for i < len(nums) {
		if target <= nums[i] {
			break
		}else if i+1 < len(nums) && target < nums[i+1] {
			i = i + 1
			break
		}else {
			i++
		}
	}
	return i
}
//二分查找
func searchInsert2(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var mid int
	for left <= right {
		mid = left + (right - left)  / 2
		if nums[mid] == target {
			return mid
		}else if target > nums[mid] {
			left = mid + 1
		}else if target < nums[mid] {
			right = mid - 1
		}
	}

	return left
}