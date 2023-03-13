package main

import "fmt"

// 在函数体外声明变量应使用var或const
const n = 1000

var s string

// 这种不带声明格式的只能在函数体中出现
// a, b, c := 5, 7, "abc"
func test() {
	a, b, c := 5, true, "abc"
	//（在函数体中的）声明后必须使用
	//全局声明的，不使用也可
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func main() {
	//多变量声明，自动推断类型
	a, b, c := 5, true, "abc"
	//（在函数体中的）声明后必须使用
	//全局声明的，不使用也可
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
