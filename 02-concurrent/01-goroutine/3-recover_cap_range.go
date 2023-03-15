/*
recover捕获范围
用recover捕获异常时，只能捕获当前goroutine的panic，不能捕获其他goroutine发生的panic。
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() { //捕获panic
		if e := recover(); e != nil {
			fmt.Printf("main recover:%v\n", e)
		}
	}()

	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Printf("sub recover:%v\n", e)
			}
		}()
		panic("sub func panic!!!")
		fmt.Println("111") // 发生panic后，不会打印
	}()

	panic("main func panic!!!")
	fmt.Println("222") // 发生panic后，不会打印
	time.Sleep(time.Second)
}
