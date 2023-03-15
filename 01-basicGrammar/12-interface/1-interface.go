/*
声明：指定了接收者的函数。
接口：含有【一组方法】的声明。
实现：当某一种类型实现了【所有】这些声明的方法，那么就称这种类型为该接口的一种实现。

接口-类型变量创建：（两种方法）某类型已实现接口的所有方法

	（1）直接创建该类型变量，即可调用对应接口方法
	（2）先创建接口变量，再将接口变量（new）接入到具体的某类型（注意：得到的是指针）
*/
package main

import "fmt"

type Phone interface {
	Call()
	SendMessage()
}

type Apple struct {
	PhoneName string
}

func (a Apple) Call() {
	fmt.Printf("%s有打电话功能\n", a.PhoneName)
}

func (a Apple) SendMessage() {
	fmt.Printf("%s有发短信功能\n", a.PhoneName)
}

type HuaWei struct {
	PhoneName string
}

func (h HuaWei) Call() {
	fmt.Printf("%s有打电话功能\n", h.PhoneName)
}

func (h HuaWei) SendMessage() {
	fmt.Printf("%s有发短信功能\n", h.PhoneName)
}

func main() {
	//对于Phone接口，已经给出了两种类型实现（Apple、Huawei）
	//第一种：直接声明某类型变量，用该类型变量调用所实现的方法
	a := Apple{"apple"}
	b := HuaWei{"huawei"}
	a.Call()
	a.SendMessage()
	b.Call()
	b.SendMessage()
	fmt.Printf("%#v\n", a) //main.Apple{PhoneName:"apple"}

	//第二种：先声明接口变量，将该接口变量创建成某类已实现的类型指针
	var phone Phone                    // 声明一个接口类型phone
	phone = new(Apple)                 // 注意这种创建方式，new函数参数是接口的实现
	fmt.Printf("%#v\n", phone)         //&main.Apple{PhoneName:""}
	phone.(*Apple).PhoneName = "Apple" // 这里使用断言给phone的成员赋值，后面会讲到接口的断言
	phone.Call()
	phone.SendMessage()
	fmt.Printf("%#v\n", phone) //&main.Apple{PhoneName:"Apple"}
}
