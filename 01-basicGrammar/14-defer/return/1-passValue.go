/*
直接记住defer传入的参数固定不变，注意区分传进去的是值还是地址
*/
package main

import "fmt"

func deferRun1() {
	var num = 1
	numptr := &num
	defer fmt.Println("defer1:\n", numptr, num, *numptr) // 0xc00000e0b8 1 1
	defer func() {
		fmt.Println("defer2:\n", numptr, num, *numptr) //0xc00000e0b8 2 2
	}() //无参数
	defer func(num int, numptr *int) {
		fmt.Println("defer3:\n", numptr, num, *numptr) //0xc00000e0b8 1 2
	}(num, numptr) //参数
	/*	//等价于情况1
		defer func(num int, num2 int) {
			fmt.Println("defer4:\n", &numptr, num, num2) //0xc00000e0b8 1 1
		}(num, *numptr)
	*/

	num = 2
	fmt.Println(numptr, num, *numptr) //0xc00000e0b8 2 2
	return
}

func main() {
	deferRun1()
}
