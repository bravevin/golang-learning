package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	co := sync.NewCond(&sync.Mutex{})
	go func() {
		co.L.Lock()
		fmt.Println("lock1")
		co.Wait()
		fmt.Println("unlock1")
		co.L.Unlock()
	}()

	go func() {
		co.L.Lock()
		fmt.Println("lock2")
		co.Wait()
		fmt.Println("unlock2")
		co.L.Unlock()
	}()

	time.Sleep(2 * time.Second)
	//co.Broadcast() // 全部解锁
	co.Signal() // 单个解锁
	time.Sleep(2 * time.Second)
	co.Signal()
	time.Sleep(2 * time.Second)
}
