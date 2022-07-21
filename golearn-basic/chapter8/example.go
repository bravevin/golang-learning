package main

import (
	"fmt"
)

func main() {
	// 指针和地址
	var a int
	a = 100
	fmt.Println(a)
	var b *int
	b = &a
	*b = 999
	fmt.Println(a, b)
	fmt.Println(a == *b, b == &a)

	// 数组指针
	var arr [4]int = [4]int{3, 5, 8, 9}
	var arrP *[4]int
	arrP = &arr
	fmt.Println(arr, arrP)
	// 指针数组
	var arrPP [3]*string
	var str1 string = "123"
	var str2 string = "embed"
	var str3 string = "abc"
	arrPP = [3]*string{&str1, &str2, &str3}
	fmt.Println(arrPP)

	*arrPP[1] = "audio"
	fmt.Println(str2)

}
