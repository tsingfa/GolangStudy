/*
绑定recover创建goroutine

将协程的逻辑封装成函数，绑定recover，以此来创建goroutine
避免频繁地写defer recover函数来捕获异常
*/
package main

import (
	"fmt"
	"sync"
)

func withGoroutine(opts ...func() error) (err error) {
	var wg sync.WaitGroup
	for _, opt := range opts {
		wg.Add(1)
		// 开启goroutine，做并行处理
		go func(handler func() error) {
			defer func() {
				// 协程内部捕获panic
				if e := recover(); e != nil {
					fmt.Printf("recover:%v\n", e)
				}
				wg.Done()
			}()
			fmt.Println("hint0!!") //在发生panic之前，可以打印
			e := handler()         // 真正的逻辑调用（发生panic）
			fmt.Println("hint1!!") //在panic所在函数的同一层，无法打印

			// 取第一个报错的handler调用的错误返回（但会被跳过）
			if err == nil && e != nil {
				err = e
			}
		}(opt) // 将goroutine的函数逻辑通过封装成的函数变量传入
	}
	fmt.Println("hint2!!") //在panic所在函数的外层，可以打印
	wg.Wait()              // 等待所有的协程执行完
	return err
}

func main() {
	handler1 := func() error {
		panic("handler1 fail ")
		return nil
	}

	handler2 := func() error {
		panic("handler2 fail")
		return nil
	}

	err := withGoroutine(handler1, handler2) // 并发执行handler1和handler2两个任务，返回第一个报错的任务错误
	fmt.Println("hint3!!")                   //在panic所在函数的外层，可以打印
	fmt.Println("err:", err)                 //err赋值被跳过，传不出来
	if err != nil {
		fmt.Printf("err is:%v", err)
	}
}
