package main

import "fmt"

/**
「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：
1.     1
2.     11
3.     21
4.     1211
5.     111221
6.		312211
7.		13112221
8.		1113213211
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。
注意：整数序列中的每一项将表示为一个字符串。
 */

func main() {
	n := 5
	fmt.Println(countAndSay(n))
}

func countAndSay(n int) string {
	if n < 1 || n > 30 {
		return ""
	}
	i := 1
	originalStr := "1"
	result := ""

	return countAndSay2(n,result)
}

func countAndSay2(n int,str string) string {
	c := str[0]
	count := 0
	for i:=0;i<len(str);i++{
		if str[i] == c {
			count++
			c = str[i]
		}else {
			result += fmt.Sprintf("%v",count) + string(c)
		}
	}
	result += fmt.Sprintf("%v",count) + string(c)
	return result
}