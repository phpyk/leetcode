package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "{([])}"
	fmt.Println(isValid(str))
}

func isValid(s string) bool {
	var parentheses = []string{"()","[]","{}"}
	for s != "" && hasParenthese(s,parentheses) {
		for _,v := range parentheses {
			s = strings.ReplaceAll(s,v,"")
		}
	}
	return s == ""
}

func hasParenthese(s string, parentheses []string) bool {
	for _,v := range parentheses {
		if strings.Index(s,v) >= 0 {
			return true
		}
	}
	return false
}
