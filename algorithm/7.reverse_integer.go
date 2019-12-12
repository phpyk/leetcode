package main

import (
	"fmt"
	"strconv"
)

/**
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
示例 1:
输入: 123
输出: 321

示例 2:
输入: -123
输出: -321

示例 3:
输入: 120
输出: 21
注意:
假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
 */
func main() {
	source := -12345678
	target := reverse2(source)
	fmt.Println(target)
}
// 先转换为字符串，字符串再转为rune slice，交换字符顺序：第一个和最后一个交换，第二个和倒数第二个交换，以此类推，最后转为int
func reverse(x int) int {
	var str string
	if x < 0 {
		str = strconv.Itoa(0-x)
	}else {
		str = strconv.Itoa(x)
	}
	r := []rune(str)
	l := len(r)
	var halfLen int
	if l% 2 == 0{
		halfLen = l /2
	}else {
		halfLen = (l-1) /2
	}
	for i,_ := range r {
		if i == halfLen {
			break
		}
		tmp := r[i]
		r[i] = r[l-i-1]
		r[l-i-1] = tmp
	}
	res := string(r)
	if x < 0 {
		res = "-"+res
	}
	intRes,err := strconv.Atoi(res)
	if err != nil{
		return 0
	}
	if intRes > 1 << 31 -1 || intRes < -1 << 31 {
		return 0
	}
	return intRes
}

func reverse2(x int) int {
	ret := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		ret = ret * 10 + pop
	}
	if ret > 1 << 31 -1 || ret < -1 << 31 {
		return 0
	}
	return ret
}