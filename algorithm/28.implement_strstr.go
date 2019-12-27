package main

import (
	"fmt"
)

/**
实现 strStr() 函数。
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:
输入: haystack = "hello", needle = "ll"
输出: 2

示例 2:
输入: haystack = "aaaaa", needle = "bba"
输出: -1
 */
func main() {
	haystack := "abcdefg"
	//fmt.Println(haystack[1:5])
	//os.Exit(200)
	needle := "ef"
	fmt.Println(strStr(haystack,needle))
}

func strStr(haystack string, needle string) int {
	l := len(needle)
	if l == 0 {
		return 0
	}else {
		hasFind := false
		startIdx := 0
		endIdx := startIdx + l
		hayLen := len(haystack)
		for endIdx <= hayLen {
			if haystack[startIdx:endIdx] == needle {
				hasFind = true
				break
			}
			startIdx++
			endIdx++
		}

		if hasFind {
			return startIdx
		}else {
			return -1
		}
	}
}
