package main

import "fmt"

type User struct {
	Name string
	Age  int
	Sex  bool
}

type Student struct {
	Class string
	User
}

func (u User) sayName(name string) {
	fmt.Println("我的名字叫做", name)
}

func main() {
	s := Student{
		Class: "三年二班",
		User:  User{"miaoqi", 19, false},
	}
	u := User{
		Name: "qimiao",
		Age:  18,
		Sex:  true,
	}
	check(s)
	check(u)
}

// assertion断言
func check(v interface{}) {
	switch v.(type) { // 获取v的结构体类型
	case User:
		fmt.Println(v.(User).Name) // v.(User)断言成User类型打印Name
		fmt.Println("我是User")
	case Student:
		fmt.Println(v.(Student).Class) // v.(Student)断言成Student类型打印Class
		fmt.Println("我是Student")
	}

}
