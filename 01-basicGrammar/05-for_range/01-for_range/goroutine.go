package main

import (
	"fmt"
	"time"
)

func main() {
	var m = []int{1, 2, 3} //输出的i是0,1,2	m[i]为1,2,3

	for i := range m {
		go func() {
			fmt.Print(m[i]) //结果可能为：333
			//因为很有可能当for循环执行完之后，goroutine才开始执行，
			//这个时候val的值是指向了切片中最后一个元素。
			//所以三个goroutine打印出来的结果相同
		}()
	}

	/*
		//修改方法一:以传参方式传入
		for i := range m {
			go func(i int) {	//将i的值指向一个传参（拷贝），则不会受range_i的变化影响
				fmt.Print(m[i]) //312
			}(i)

		}
	*/
	/*
		//修改方法一:使用局部变量拷贝
		for i := range m {
			j := i
			go func() {
				fmt.Print(m[j]) //123
			}()
		}

	*/
	//阻塞1分钟等待所有goroutine运行完
	time.Sleep(time.Millisecond)
}
