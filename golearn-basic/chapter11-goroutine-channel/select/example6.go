package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	ch1 <- 1
	//ch2 <- 2
	//ch3 <- 3
	// select 可以执行可读取的channel, 哪个有读哪个
	// 如果均可读取,随机执行任意的case
	select {
	case <-ch1:
		fmt.Println("ch1")
	case <-ch2:
		fmt.Println("ch2")
	case <-ch3:
		fmt.Println("ch3")
	default:
		fmt.Println("都不满足")
	}
}
