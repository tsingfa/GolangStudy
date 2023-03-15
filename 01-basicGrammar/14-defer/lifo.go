/*
defer：方法或函数的延迟调用
调用时间：

	（1）函数return
	（2）发生panic

语法规则：

	（1）defer关键字后接函数或方法调用
	（2）延迟内容不能被括号括起来

多个defer的执行顺序：LIFO（后进先出）
*/
package main

import "fmt"

func defer1() {
	fmt.Println("defer1")
}

func defer2() {
	fmt.Println("defer2")
}

func defer3() {
	fmt.Println("defer3")
}

func main() {
	//先调用的，后执行
	defer defer1()
	defer defer2()
	defer defer3()
}
