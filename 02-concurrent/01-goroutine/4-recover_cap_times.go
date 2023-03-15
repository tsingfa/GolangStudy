/*
recover捕捉次数
一个recover只能捕获一次panic，且一一对应
*/
package main

import (
	"fmt"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("recover:%v\n", e)
		}
	}()
	panic("panic1") //recover捕获到第一个panic之后，该函数就退出了
	panic("panic2")
	fmt.Println("111") // 发生panic，不会打印
}
