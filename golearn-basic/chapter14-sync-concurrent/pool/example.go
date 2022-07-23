package main

import (
	"fmt"
	"sync"
)

func main() {
	p := &sync.Pool{}
	p.Put(1)
	p.Put(3)
	p.Put(4)
	p.Put(7)
	p.Put(9)
	for i := 0; i < 10; i++ {
		fmt.Println(p.Get())
	}
}
