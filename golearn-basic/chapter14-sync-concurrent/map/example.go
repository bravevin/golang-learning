package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//test1()
	//test2()
	test3()
}

func test1() {
	m := &sync.Map{}
	go func() {
		for {
			m.Store(1, 1)
		}
	}()
	go func() {
		for {
			fmt.Println(m.Load(1))
		}
	}()

	time.Sleep(100)
}

func test2() {
	m := &sync.Map{}
	m.Store(1, 2)
	m.LoadOrStore(3, 4)
	m.Delete(1)
	fmt.Println(m.Load(1))
	fmt.Println(m.Load(3))
}

func test3() {
	m1 := &sync.Map{}
	m1.Store(1, 2)
	m1.Store(2, 5)
	m1.Store(3, 8)
	m1.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		time.Sleep(1 * time.Second)
		//return false // false中断执行,true继续执行
		return true
	})
}
