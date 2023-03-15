/*
接口嵌套：一个接口中包含了其他接口
如果要实现外部接口，则需要实现内部嵌套的接口对应的所有方法。
*/
package main

import "fmt"

type A interface {
	run1()
}

type B interface {
	run2()
}

// 定义嵌套接口C
type C interface {
	A
	B
	run3()
}

type Runner struct{}

// 实现嵌套接口A的方法
func (r Runner) run1() {
	fmt.Println("run1!!!!")
}

// 实现嵌套接口B的方法
func (r Runner) run2() {
	fmt.Println("run2!!!!")
}

func (r Runner) run3() {
	fmt.Println("run3!!!!")
}

func main() {
	var runner C
	runner = new(Runner) // runner实现了C接口的所有方法
	runner.run1()
	runner.run2()
	runner.run3()
}
