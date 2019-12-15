package main

import (
	"fmt"
)

func main() {
	var nums  = []int{5,5,5,4,3,2,2,1,1}
	length := removeDuplicates(nums)
	fmt.Println(length)
}
//删除数组中的重复项目
//思路：把非重复项排序，剩余的重复项丢弃
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j:= 1;j<len(nums);j++{
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	nums = nums[:i+1]
	return i+1
}
