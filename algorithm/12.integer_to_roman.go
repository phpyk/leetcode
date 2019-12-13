package main

import "fmt"

/**
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-to-roman
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func main() {
	n := 999
	res := intToRoman(n)
	fmt.Println(res)
}

func intToRoman(n int) string {
	if n > 3999 {
		return ""
	}
	var intArr = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1,}
	var strArr = []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
	//var valMap = map[int]string{1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL", 50: "L", 90: "XC", 100: "C", 400: "CD", 500: "D", 900: "CM", 1000: "M",}
	var result string
	for i,v := range intArr {
		for n >= v {
			n -= v
			result += strArr[i]
		}
	}
	return result
}