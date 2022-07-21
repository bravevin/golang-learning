package main

import "fmt"

func main() {
	// 数组 array
	var arrExample [3]int
	arrExample = [3]int{1, 3, 5}
	fmt.Println(arrExample)

	arr1 := [3]int{2, 3, 4}          // 隐式声明
	arr2 := [...]int{1, 3, 5, 9, 13} //不定长数组
	arr3 := new([10]int)
	arr3[6] = 10
	fmt.Println(arr1, arr2, arr3)
	arr4 := [...]string{"dog", "cat", "monkey"}
	for i := 0; i < len(arr4); i++ {
		fmt.Println(arr4[i] + "RUN")
	}
	fmt.Println(len(arr4) == cap(arr4))
	// arr4[3] = "I love to run";

	// 二维数组
	er := [3][3]int{
		{1, 3, 5},
		{3, 5, 7},
		{4, 6, 9},
	}
	fmt.Println(er)

	newArr := [3]int{12, 32, 56}
	// 切片 slice
	// 切片前闭后开原则
	s1 := newArr[:] // 截取全部
	fmt.Println(s1)
	s2 := newArr[:2] // 从头开始到索引1
	fmt.Println(s2)
	// 测试
	slice1 := newArr[1:] // 截取到结尾
	slice1[0] = 13       // 修改切片会影响原数组
	fmt.Println(newArr, slice1)

	slice2 := newArr[2:]       // 截取到结尾
	slice2 = append(slice2, 5) // 返回一个新切片
	fmt.Println(newArr, slice2)
	// copy函数
	arrOne := [3]int{0, 1, 2}
	cluster1 := arrOne[:]  // [0,1,2]
	cluster2 := arrOne[2:] // [2]
	cluster3 := arrOne[:2] // [0,1]
	fmt.Println(cluster1, cluster2, cluster3)
	cluster1 = append(cluster1, 5) // [0,1,2,5]
	//copy(cluster1, cluster2) // [2,1,2,5] // 默认从头部copy
	//copy(cluster1[1:], cluster2) // [0,2,2,5] //start|end 前闭后开
	copy(cluster1[2:], cluster3) // [0,1,0,1] //start|end 前闭后开
	fmt.Println(cluster1, cluster2)
	fmt.Println(cluster1, cluster3)
}
