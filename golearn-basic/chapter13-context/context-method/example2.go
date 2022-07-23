package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1.WithCancel clear表示结束
	//ctx, clear := context.WithCancel(context.Background())

	// 2.WithValue ctx包含所传的参数
	//ctx1 := context.WithValue(context.Background(), "name", "qimiao")
	//ctx, clear := context.WithCancel(ctx1)

	// 3.WithDeadline ddl表示ctx截止的时间
	//ctx, clear := context.WithDeadline(context.Background(), time.Now().Add(time.Second*4))

	// 4.WithTimeout timeout从现在开始开启固定时间的计时 WithDeadline变形
	ctx, clear := context.WithTimeout(context.Background(), time.Second*4)

	message := make(chan int)
	go son(ctx, message)
	for i := 0; i < 10; i++ {
		message <- i
	}
	clear()
	time.Sleep(time.Second)
	fmt.Println("主进程结束")
}

func son(ctx context.Context, msg chan int) {
	t := time.Tick(time.Second)
	for _ = range t {
		select {
		case m := <-msg:
			fmt.Printf("接收到值%d\n", m)
		case <-ctx.Done():
			fmt.Println("我结束了")
			//fmt.Println("我结束了", ctx.Value("name")) // 获取WithValue所传的参数
			return
		}
	}
}
