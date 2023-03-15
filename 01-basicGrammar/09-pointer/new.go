package main

import "fmt"

func main() {
	s := new(string)
	fmt.Printf("赋值前s: %v\n", s)     //有分配内存地址
	fmt.Printf("赋值前s的内容: %v\n", *s) //内容为空串

	*s = "Golang学习之路"
	fmt.Printf("赋值后s: %v\n", s)
	fmt.Printf("赋值后s的内容: %v\n", *s)
}
