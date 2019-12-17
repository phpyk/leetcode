package main

import "fmt"

func main() {
	var nums  = []int{5,5,5,5,4,3,2,5,2,7,1,1}
	var val = 1
	fmt.Printf("%p \n",nums)
	fmt.Println("length =",removeElement(nums,val))
}

func removeElement(nums []int, val int) int {
	fmt.Printf("%p \n",nums)
	for i:=0;i<len(nums); {
		if nums[i] == val {
			nums = append(nums[0:i],nums[i+1:]...)
			fmt.Printf("%p \n",nums)
		}else {
			i++
		}
	}
	return len(nums)
}
