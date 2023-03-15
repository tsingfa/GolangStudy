/*
空接口：没有任何方法声明的接口
所有的类型都实现了空接口，因此空接口可以存储任意类型的数值。
以空接口作为参数，表示可接受任意类型的参数，如fmt.Print系列函数
*/
package main

import (
	"fmt"
)

func main() {
	var any interface{} //空接口：没有任何方法声明的接口
	any = 10
	fmt.Println(any)

	any = "golang"
	fmt.Println(any)

	any = true
	fmt.Println(any)
}
