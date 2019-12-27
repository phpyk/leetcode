package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func main() {

	fmt.Println(strings.Index("",""))
	os.Exit(1)
	//声明并初始化
	ch := make(chan int,3)
	ch <- 2
	ch <- 1
	ch <- 3
	elem1 := <- ch
	fmt.Println("first element is :",elem1)
	//获取通道长度
	fmt.Println("length is :", len(ch))
	fmt.Println("second element is :",<-ch)
	fmt.Println("third element is :",<-ch)
	fmt.Println("length is :", len(ch))
	runtime.GOMAXPROCS(5)
}
