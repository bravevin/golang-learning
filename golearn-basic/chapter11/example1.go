package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup //goroutine管理器
	wg.Add(1)
	go Run(&wg)
	wg.Wait()

	//time.Sleep(time.Second * 1)
	//for i := 1; i < 8; i++ {
	//	fmt.Println(i)
	//}
}

func Run(wg *sync.WaitGroup) {
	fmt.Println("开始跑...")
	wg.Done()
}
