/*
指针运算符
&a 对变量取址
*ptr,对指针取值
*/

package main

import "fmt"

func main() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)

	/*  & 和 * 运算符实例 */
	ptr = &a                          /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("%-7s:%v\n", "a", a)   //4
	fmt.Printf("%-7s:%v\n", "&a", &a) //0xc00000e0b8
	//fmt.Printf("%-7s:%v\n", "*a", *a)		//对该变量的间接引用无效
	fmt.Printf("%-7s:%v\n", "*(&a)", *(&a))     //4
	fmt.Printf("%-7s:%v\n", "ptr", ptr)         //0xc00000e0b8	//指针指向的地址
	fmt.Printf("%-7s:%v\n", "&ptr", &ptr)       //0xc00000a028	//指针本身的地址
	fmt.Printf("%-7s:%v\n", "*ptr", *ptr)       //4
	fmt.Printf("%-7s:%v\n", "&(*ptr)", &(*ptr)) //0xc00000e0b8
}
