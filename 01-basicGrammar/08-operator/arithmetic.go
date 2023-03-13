/*
算术运算符
+,-,*,/,%,++,--
加、减、乘、除、求余、自增、自减
*/
package main

import "fmt"

func main() {

	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c) //31
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c) //11
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c) //210
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c) //2（仅保留整数部分，向下取整）
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c) //1
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a) //22
	a = 21                            // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a) //20
}
