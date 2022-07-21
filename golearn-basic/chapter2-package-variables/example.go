package main

import (
	"fmt"
	"golearn-basic/testpackage"
)

func main() {
	var a string = "hello 2020"
	b := "hello 1010"
	fmt.Println(a, b, "hello world")
	fmt.Println(testpackage.Abc)
	//fmt.Println(testpackage.ace) // 获取不到私有变量
	fmt.Println(testpackage.B)
	fmt.Println(testpackage.Test())
}
