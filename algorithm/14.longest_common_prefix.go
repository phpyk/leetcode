package main

import "fmt"

/**
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
示例 1:
输入: ["flower","flow","flight"]
输出: "fl"
示例 2:
输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:
所有输入只包含小写字母 a-z 。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func main() {
	//strArr := []string{"flower", "flow", "flight",}
	strArr := []string{"aa","a"}
	fmt.Println(longestCommonPrefix(strArr))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	firstStr := strs[0]
	commonPrefix := ""
	var runeArr [][]rune
	//先将每个str转成[]rune，避免在for循环中重复转换
	for _,str := range strs {
		runeArr = append(runeArr,[]rune(str))
	}

	for i,v := range firstStr {
		if !valueEqual(i,v, runeArr) {
			break
		}
		commonPrefix += string(v)
	}
	return commonPrefix
}

func valueEqual(i int, v rune, runeArr [][]rune) bool {
	for j:=1;j<len(runeArr);j++ {
		if len(runeArr[j])-1 < i || v != runeArr[j][i] {
			return false
		}
	}
	return true
}