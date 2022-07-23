package main

import "fmt"

func main() {
	chan1 := make(chan int)
	var reader <-chan int = chan1
	var writer chan<- int = chan1

	go SetChan(writer)

	GetChan(reader)
}

func SetChan(c1 chan<- int) {
	for i := 0; i < 10; i++ {
		c1 <- i
		fmt.Printf("set函数存%d\n", i)
	}
}

func GetChan(c2 <-chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("get函数读%d\n", <-c2)
	}
}
