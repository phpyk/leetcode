package main

import "fmt"

/**
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
示例 1:
输入: 121
输出: true

示例 2:
输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

示例 3:
输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main() {
	a := 12321
	fmt.Println(isPalindrome2(a))
}
//将数字的每一位放到int slice中，遍历slice，比较每一位与其对称位的值是否相对
//时间复杂度O(n)，空间复杂度O(n)
func isPalindrome(x int) bool {
	if x < 0 || (x > 0 && x % 10 == 0) {
		return false
	}
	var nums []int
	for x > 0 {
		y := x % 10
		nums = append(nums,y)
		x /= 10
	}
	l := len(nums)
	var halfLen int
	//只遍历slice的一半即可
	if l% 2 == 0{
		halfLen = l /2
	}else {
		halfLen = (l-1) /2
	}
	var isPalindrome bool = true
	for i,v := range nums {
		if i == halfLen {
			break
		}
		if v != nums[l-i-1] {
			isPalindrome = false
			break
		}
	}
	return isPalindrome
}

//最优解
//时间复杂度O(n)，空间复杂度O(1)
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	r := 0
	tmp := x
	for tmp > 0 {
		r = r*10 + tmp % 10
		tmp /= 10
	}
	if r == x {
		return true
	}
	return false
}
