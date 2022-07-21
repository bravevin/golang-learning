package testpackage

import "fmt"

func Test() (ret string) {

	// 同包下的go文件中的变量是可以互相访问的
	fmt.Println(B, mallUserSevice, Abc, ace)
	ret = B + mallUserSevice + Abc + ace
	return ret
}
