/*
断言
可将任意类型的变量赋给空接口，但是反过来不行（除非断言）
*/
package main

import "fmt"

func main() {
	var a int = 1
	var i interface{} = a //可将任意类型的变量赋给空接口
	// 但是反过来不行（除非断言）
	//var b int = i	//无法将i(interface{})用作int
	var b int
	var ok bool
	b, ok = i.(int)
	fmt.Println(b, ok)
	//若没有bool值判断，而当前动态类型与断言类型不一致，则会报panic
	//c := i.(string) //panic: interface conversion: interface {} is int, not string
	//fmt.Println(c)
	c, ook := i.(string)
	fmt.Printf("%#v,%v", c, ook) //"",false	//无法赋值，故c为对应零值（空串）

}
