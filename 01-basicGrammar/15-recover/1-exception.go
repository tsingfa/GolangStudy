/*
recover异常捕获
*/
package main

import "fmt"

func main() {
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("出现了panic,使用reover获取信息:", error)
		}
	}()
	fmt.Println("11111111111")
	panic("出现panic") //有recover之后，程序不会在panic处中断，panic信息会被recover捕获
	fmt.Println("22222222222")
}
