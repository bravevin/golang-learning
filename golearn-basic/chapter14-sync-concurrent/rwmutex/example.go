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
	l := &sync.RWMutex{} // 读写锁
	go ReadLockFun(l)
	go ReadLockFun(l)
	go ReadLockFun(l)
	go ReadLockFun(l)
	go LockFun(l)
	go LockFun(l)

	for {
	} //起到阻塞作用,显示goroutine执行过程
}

// LockFun 增加互斥锁,一旦lock住,其他goroutine函数无法执行
func LockFun(lock *sync.RWMutex) {
	lock.Lock() // 写的时候,排斥写锁和读锁
	fmt.Println("疯狂刮痧")
	time.Sleep(1 * time.Second)
	lock.Unlock()
}

func ReadLockFun(readLock *sync.RWMutex) {
	readLock.RLock() //读的时候,排斥写锁而不排斥读锁
	fmt.Println("疯狂治疗")
	time.Sleep(1 * time.Second)
	readLock.RUnlock()
}
