/*
main()对应创建有主协程，主协程过早退出使得其他协程来不及执行
*/
package main

import (
	"fmt"
	"time"
)

func myGroutine1() {
	fmt.Println("myGroutine")
}

func main() {
	go myGroutine1()
	fmt.Println("end!!!")
	time.Sleep(time.Second) //等待，防止主协程过早退出
}
