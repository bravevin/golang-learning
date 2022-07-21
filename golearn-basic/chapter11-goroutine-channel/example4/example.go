package main

import "fmt"

func main() {
	var c1 = make(chan int, 5)
	c1 <- 1
	c1 <- 2
	c1 <- 3
	//close(c1) // 1. uncomment this = panic: send on closed channel
	c1 <- 4
	//c1 <- 5     // 2.comment 1&&2 fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println(<-c1)

}
