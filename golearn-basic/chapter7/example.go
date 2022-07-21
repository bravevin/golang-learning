package main

import "fmt"

func main() {
	r1, r2 := a(443, "I am there for you")
	fmt.Println(r1, r2)
	b(56, "1", "2", "3", "4")

	c := func(expr string) {
		fmt.Println(expr)
	}
	c("hello in anonymous func")

	strArr := []string{"I", "am", "a", "string"}
	b(9607, strArr...)

	(func() {
		fmt.Println("I am an instant func")
	})()

	mo()(4)

	defer delay()
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
}

func delay() {
	fmt.Println("我是延迟调用的函数")
}

func mo() func(args int) {
	return func(args int) {
		fmt.Println("闭包函数", args)
	}
}

func b(data1 int, data2 ...string) {
	fmt.Println(data1, data2)
}

func a(data1 int, data2 string) (ret1 int, ret2 string) {
	ret1 = data1
	ret2 = data2
	return // 隐式返回ret1和ret2
}
