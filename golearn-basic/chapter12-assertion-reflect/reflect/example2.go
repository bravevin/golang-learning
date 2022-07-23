package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
	Sex  bool
}

type Student struct {
	Class string
	User
}

func (u User) SayName(name string) { // 必须大写开头
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
	//check(u)
	s1 := Student{"三年四班", u}
	check(s1)

	elem(&s) // 修改原始数据必须传递引用地址
	fmt.Println(s)

	method(u)
}

// reflect反射
func check(inter interface{}) {
	t := reflect.TypeOf(inter)  // 动态获取输入参数接口中值的类型
	v := reflect.ValueOf(inter) // 动态获取输入参数接口中数据的值
	//fmt.Println(t, v)

	for i := 0; i < t.NumField(); i++ {
		// reflect.ValueOf(inter).Field(index) 获取对应索引下的值
		//fmt.Println(v.Field(i))
	}
	fmt.Println(v.FieldByName("Class"))      // 根据字段名获取字段值
	fmt.Println(v.FieldByIndex([]int{1, 0})) // 数组依次存放层级的索引

	ty := t.Kind() // reflect.TypeOf(inter).Kind() // 获取inter具体的数据类型

	if ty == reflect.Struct {
		fmt.Println("struct")
	}

	if ty == reflect.String {
		fmt.Println("string")
	}

	if ty == reflect.Int {
		fmt.Println("int")
	}
}

func elem(inter2 interface{}) {
	v := reflect.ValueOf(inter2)
	// v.Elem() 获取原始数据并操作
	v.Elem().FieldByName("Class").SetString("四年三班")
	v.Elem().FieldByName("User").Field(1).SetInt(26)
	v.Elem().FieldByIndex([]int{1, 2}).SetBool(true)
	fmt.Println(inter2)
}

func method(inter3 interface{}) {
	v := reflect.ValueOf(inter3)
	// 获取method方法并执行
	m := v.Method(0)
	fmt.Println(m)
	m.Call([]reflect.Value{reflect.ValueOf("vincent")})
}
