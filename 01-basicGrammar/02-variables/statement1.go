package main

import "fmt"

func main() {
	//未初始化声明（零值）
	var i int
	var f float64
	var b bool
	var s string
	//以下输出使用左对齐15个字符长度
	fmt.Printf("%-15s %#v\n", "default int:", i)
	fmt.Printf("%-15s %#v\n", "default float:", f)
	fmt.Printf("%-15s %#v\n", "default bool:", b)
	fmt.Printf("%-15s %#v\n\n", "default string:", s)

	var a1 *int
	var a2 []int
	var a3 map[string]int
	var a4 chan int
	var a5 func(string) int //函数变量，规定了参数和返回值类型
	var a6 error            // error 是接口
	fmt.Printf("%-15s %#v\n", "default *int:", a1)
	fmt.Printf("%-15s %#v\n", "default []int:", a2)
	fmt.Printf("%-15s %#v\n", "default map:", a3)
	fmt.Printf("%-15s %#v\n", "default chan:", a4)
	fmt.Printf("%-15s %#v\n", "default func:", a5)
	fmt.Printf("%-15s %#v\n", "default error:", a6)

}
