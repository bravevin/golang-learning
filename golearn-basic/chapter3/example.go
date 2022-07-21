package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 uint
	var num2 int
	num1 = 999
	num2 = -999
	var num3 float64
	num3 = 3.1415926
	var string1 string
	string1 = "这样就是字符串类型"
	var boolean bool
	boolean = true
	var number uint
	number = 123
	fmt.Println(num1, num2, num3, string1, boolean)
	fmt.Printf("%T\n", number) // 展示数据类型
	var string2 string = "12315"
	val, _ := strconv.Atoi(string2)
	fmt.Println(val)
	// 隐式声明
	val1 := 12813
	val2 := "hello world"
	val3 := 3.1415928
	val4 := false
	fmt.Printf("%T %T %T %T\n", val1, val2, val3, val4)
}
