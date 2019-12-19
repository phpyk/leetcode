package main

import (
	"fmt"
)

func main() {
	m := 1
	n := 1
	fmt.Printf("%p\n",&m)
	fmt.Printf("%p\n",&n)

	a := [5]int{1,2,3,}
	b := [5]int{1,2,3,}

	fmt.Printf("%p\n",&a)
	fmt.Printf("%p\n",&b)

}
