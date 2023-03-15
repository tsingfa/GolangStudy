/*
recover和panic可以总结为以下两点：
 1. recover()只能恢复当前函数级或
    以当前函数为首的调用链中的函数中的panic(),
    恢复后调用当前函数结束,但是调用此函数的函数继续执行
 2. 函数发生了panic之后会一直向上传递,
    如果直至main函数都没有recover()，程序将终止，
    如果是碰见了recover()，将被recover捕获。
*/
package main

import "fmt"

func testPanic1() {
	fmt.Println("testPanic1上半部分")
	testPanic2()
	fmt.Println("testPanic1下半部分")
}

func testPanic2() {
	defer func() {
		recover()
	}()
	fmt.Println("testPanic2上半部分")
	testPanic3() //在该函数级中panic被recover捕获，当前函数中止，上级函数继续
	fmt.Println("testPanic2下半部分")
}

func testPanic3() {
	fmt.Println("testPanic3上半部分")
	panic("在testPanic3出现了panic")
	fmt.Println("testPanic3下半部分")
}

func main() {
	fmt.Println("程序开始")
	testPanic1()
	fmt.Println("程序结束")
}
