package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//1.sleep() 添加WaitGroup之前

	wg := &sync.WaitGroup{}
	wg.Add(2) // 表示需要处理2个goroutine

	go func() {
		time.Sleep(8 * time.Second)
		fmt.Println("hp-1")
		wg.Done()
		// 标记该goroutine结束delta-1
	}()
	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("hp-1")
		wg.Done()
		// 标记该goroutine结束delta-1
	}()
	wg.Wait()
	fmt.Println("hp=0 game over")
}

func sleep() {
	go func() {
		time.Sleep(8 * time.Second)
		fmt.Println(8)
	}()
	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println(6)
	}()

	time.Sleep(6 * time.Second) //goroutine函数可能不会执行
}
