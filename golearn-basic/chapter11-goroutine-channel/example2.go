package main

import "fmt"

func main() {
	c1 := make(chan int) // 无缓冲区
	go func() {
		c1 <- 1 // 写入1
	}()
	fmt.Println(<-c1) // 读取1
	fmt.Println("========")
	c2 := make(chan int) // 无缓冲区

	go func() {
		for i := 0; i < 5; i++ {
			c2 <- i //debug 断点
		}
	}()

	for j := 0; j < 5; j++ {
		fmt.Println(<-c2) //debug 断点
	}

	fmt.Println("========")
	c3 := make(chan int, 5) // 有5个缓冲区

	go func() {
		for m := 0; m < 10; m++ {
			c3 <- m //debug 断点 存满5个后可再取出
		}
	}()

	for n := 0; n < 10; n++ {
		fmt.Println(<-c3) //debug 断点
	}
}
