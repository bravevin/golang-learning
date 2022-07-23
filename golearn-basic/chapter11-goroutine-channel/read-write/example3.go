package main

import "fmt"

func main() {
	var c1 = make(chan int, 5)
	var readc <-chan int = c1
	var writec chan<- int = c1
	writec <- 1
	fmt.Println(<-readc)
}
