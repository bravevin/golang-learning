package main

import "fmt"

func main() {
	// if else
	a := 10
	if a > 9 {
		fmt.Println("a > 9")
	} else {
		fmt.Println("a <= 9")
	}
	// switch
	b := 2
	switch b {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
		fallthrough // 继续向下执行
	case 3:
		fmt.Println(3)
		fallthrough // 继续向下执行
	case 4:
		fmt.Println(4)
	default:
		fmt.Println("this is default")
	}
	// for
	// form 1
	c := 1
	for {
		c++
		fmt.Println(c)
		if c > 2 {
			break
		}
	}
	// form 2
	// 完整形式
	for i := 0; i < 8; i += 3 {
		fmt.Println(i)
	}
	// form 3 相当于 while
	// 简略版
	j := 10
	for j < 20 {
		j++
		fmt.Println(j)
	}
	// for break
	for k := 1; k < 9; k++ {
		if k == 3 {
			break
		}
		fmt.Println(k)
	}
	// for continue
	for l := 2; l < 5; l++ {
		if l == 3 {
			continue
		}
		fmt.Println(l)
	}
	// goto
	number1 := 2
	for number1 < 6 {
		number1++
		if number1 == 4 {
			goto B
		}
	}
B:
	fmt.Println("a a a rush B")
}
