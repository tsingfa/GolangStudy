/*
1个指针变量可以指向任何一个值的内存地址，
它所指向的值的内存地址在 32 和 64 位机器上分别占用 4 或 8 个字节，
占用字节的大小与所指向的值的大小无关。
*/
package main

import "fmt"

func main() {
	var a *int //赋值前还是空指针nil
	fmt.Printf("赋值前a: %v\n", a)

	var b int = 3
	a = &b                          // a存放的是b的地址
	fmt.Printf("赋值后a: %v\n", a)     // 打印出b的地址
	fmt.Printf("a指向的内容是: %v\n", *a) // *a表示a存放的地址对应的内存存放的内容，这里为整数3
}
