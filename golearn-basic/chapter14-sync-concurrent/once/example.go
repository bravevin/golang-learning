package main

import (
	"fmt"
	"sync"
)

func main() {
	o := &sync.Once{}
	// o.Do的func只执行一次
	// o.done
	// 初始等于0 函数执行完成等于1,
	// 再次调用Do时不等于0,无法再次执行
	/*
		source code
		 defer atomic.StoreUint32(&o.done, 1)
		f()
	*/

	for i := 0; i < 10; i++ {
		o.Do(func() {
			fmt.Println(i)
		})
	}

	for j := 1; j < 10; j++ {
		o.Do(func() {
			fmt.Println(j)
		})
	}

	for m := 2; m < 10; m++ {
		o.Do(func() {
			fmt.Println(m)
		})
	}

}
