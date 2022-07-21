package main

import "fmt"

func delay() {
	fmt.Println("我是延迟调用的函数")
}

// 闭包 返回值为一个函数
func mo() func(args int) {
	return func(args int) {
		fmt.Println("闭包函数", args)
	}
}

// data2 可变长度的参数string data2收集成数组
func b(data1 int, data2 ...string) {
	fmt.Println(data1, data2)
}

// func 函数名(形参1 类型, 形参2 类型) (返回值1 类型,返回值2 类型)
func a(data1 int, data2 string) (ret1 int, ret2 string) {
	ret1 = data1
	ret2 = data2
	return // 隐式返回ret1和ret2
}

func main() {
	r1, r2 := a(443, "I am there for you")
	fmt.Println(r1, r2)
	b(56, "1", "2", "3", "4")

	c := func(expr string) {
		fmt.Println(expr)
	}
	c("hello in anonymous func")

	strArr := []string{"I", "am", "a", "string"}
	// strArr...展开数组strArr
	b(9607, strArr...)

	// 立即执行函数
	(func() {
		fmt.Println("I am an instant func")
	})()

	// 闭包函数执行
	mo()(4)

	defer delay() // 最后调用
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
}
