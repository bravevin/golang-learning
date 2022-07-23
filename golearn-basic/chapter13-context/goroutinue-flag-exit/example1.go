package main

import (
	"fmt"
	"time"
)

func main() {
	flag := make(chan bool)
	message := make(chan int)
	go son(flag, message)
	for i := 0; i < 10; i++ {
		message <- i
	}
	flag <- true
	time.Sleep(time.Second)
	fmt.Println("主进程结束")
}

func son(flag chan bool, msg chan int) {
	t := time.Tick(time.Second)
	for _ = range t {
		select {
		case m := <-msg:
			fmt.Printf("接收到值%d\n", m)
		case <-flag:
			fmt.Println("我结束了")
			return
		}
	}
}
