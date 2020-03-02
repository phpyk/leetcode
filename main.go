package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(strings.Index("",""))
	////声明并初始化
	//ch := make(chan int,3)
	//ch <- 2
	//ch <- 1
	//ch <- 3
	//elem1 := <- ch
	//fmt.Println("first element is :",elem1)
	////获取通道长度
	//fmt.Println("length is :", len(ch))
	//fmt.Println("second element is :",<-ch)
	//fmt.Println("third element is :",<-ch)
	//fmt.Println("length is :", len(ch))
	//runtime.GOMAXPROCS(5)

	fmt.Println("In main()")
	go longWait()
	go shortWait()
	//go longestWait()
	fmt.Println("About to sleep in main()")
	// sleep works with a Duration in nanoseconds (ns) !
	time.Sleep(15 * 1e9)
	fmt.Println("At the end of main()")
}

func longestWait() {
	fmt.Println("Beginning longestWait()")
	time.Sleep(15 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longestWait()")
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of shortWait()")
}
