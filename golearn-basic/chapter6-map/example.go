package main

import "fmt"

func main() {
	var m1 map[string]string // map形式
	m1 = map[string]string{}
	m1["name"] = "qm"
	m1["sex"] = "male"

	m2 := map[string]string{} //隐式声明
	m2["dir"] = "windows"
	m2["os"] = "win11"

	m3 := make(map[string]string) //make形式
	m3["direction"] = "west"

	fmt.Println(m1, m2, m3)

	m4 := map[int]string{} // key-int/value-string
	m4[0] = "12345"
	m4[1] = "RANDOM"
	m4[2] = "string"
	fmt.Println(m4)
	delete(m4, 1)
	fmt.Println(m4)

	m5 := map[string]interface{}{} //key-string/value-任意形式
	m5["name"] = "sb"
	m5["sex"] = "female"
	m5["age"] = 26
	m5["ad"] = "car"
	m5["eat"] = [3]int{3, 5, 6}
	delete(m5, "sex")
	for k, v := range m5 {
		fmt.Println(k, v)
	}

	m6 := map[string][]string{"sex": {"1", "2", "3"}, "address": {"bj", "sh", "gz"}}
	// key-string value-string数组
	fmt.Println(m6)
}
