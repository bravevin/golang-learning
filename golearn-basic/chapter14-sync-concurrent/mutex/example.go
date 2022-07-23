package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	SyncClass()
}

func SyncClass() {
	l := &sync.Mutex{}
	go LockFun(l)
	go LockFun(l)
	go LockFun(l)
	go LockFun(l)
	for {
	} //起到阻塞作用,显示goroutine执行过程
}

// LockFun 增加互斥锁,一旦lock住,其他goroutine函数无法执行
func LockFun(lock *sync.Mutex) {
	lock.Lock()
	fmt.Println("疯狂刮痧")
	time.Sleep(1 * time.Second)
	lock.Unlock()
}
