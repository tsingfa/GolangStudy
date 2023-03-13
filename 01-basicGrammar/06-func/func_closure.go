// 函数闭包（匿名函数）
/*
匿名函数就是一个"内联"语句或表达式。
匿名函数的优越性在于可以直接使用函数内的变量，不必声明，
这样有时候可以使代码更简单，增强代码的可读性
*/
package main

import "fmt"

func getNumber() func() int { //返回一个【以int为返回值的函数变量】
	i := 0
	fmt.Printf("hint1:i=%v\n", i)
	return func() int { //返回的int型函数变量
		i += 1
		fmt.Printf("hint2:i=%v\n", i)
		return i
	}
}

func main() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	fmt.Println("声明nextNumber函数：")
	nextNumber := getNumber() //得到所返回的int型函数，但是还没执行里面的内容，所以i=0

	fmt.Println("调用nextNumber函数：")
	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber()) //1
	fmt.Println(nextNumber()) //2
	fmt.Println(nextNumber()) //3

	/* 创建新的函数 nextNumber1，并查看结果 */
	fmt.Println("声明nextNumber1函数：")
	nextNumber1 := getNumber()
	fmt.Println("调用nextNumber1函数：")
	fmt.Println(nextNumber1()) //1
	fmt.Println(nextNumber1()) //2
}
