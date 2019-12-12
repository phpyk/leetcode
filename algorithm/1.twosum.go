package main

import "fmt"

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15,16], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
 */
func main() {
	var nums = []int{2, 7, 11, 15,18,20,33,78,10}
	target := 30
	ind := twoSum2(nums, target)
	fmt.Println(ind)
}

func twoSum(nums []int, target int) []int {
	var index []int
	var l int = len(nums)
	for i:=0;i<l;i++ {
		for j:= i+1 ;j< l;j++{
			if nums[i]+nums[j] == target{
				return append(index,i,j)
			}
		}
	}
	return index
}

func twoSum2(nums []int, target int) []int {
	var index []int
	l := len(nums)
	var diff int
	for i:=0;i<l;i++ {
		diff = target- nums[i]
		for j:= i+1 ;j< l;j++{
			if nums[j] == diff {
				return append(index,i,j)
			}
		}
	}
	return index
}