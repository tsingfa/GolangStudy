/*
defer配合recover一起处理panic
go语言中用panic来抛出异常，用recover来捕获异常，
所以当程序出现异常时，可以用defer recover来捕获异常。
*/
package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	a := 1
	b := 0
	fmt.Println("result:", a/b)
}
