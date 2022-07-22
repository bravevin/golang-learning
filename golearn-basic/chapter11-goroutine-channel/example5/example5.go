package main

import "fmt"

func main() {
	c1 := make(chan int, 5)
	c1 <- 1
	c1 <- 2
	c1 <- 3
	c1 <- 4
	c1 <- 5
	close(c1) // 必须关闭该channel
	for v := range c1 {
		fmt.Println(v)
	}
}
