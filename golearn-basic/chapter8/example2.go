package main

import "fmt"

func main() {
	var str1 string = "I am the string"
	var str2 *string = &str1
	*str2 = "change the string value"
	fmt.Println(str1, *str2)

}
