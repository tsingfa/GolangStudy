package main

import "fmt"

type Reader interface {
	Read() int
	Read_str() string
}

type MyReader struct {
	a, b int
	info string
}

func (m *MyReader) Read() int {
	return m.a + m.b
}
func (m MyReader) Read_str() string {
	return m.info
}

func DoJob(r Reader) {
	fmt.Printf("myReader is %d\n", r.Read())
	fmt.Println("info:", r.Read_str())
}

func main() {
	myReader := &MyReader{2, 5, "hello"} //结构体指针
	//该类型实现了某（些）接口，可理解为：该类型获得了（这些）接口的“通行证”
	//所以该接口参数可以传进对应实现了接口的类型（实现了方法的类型）
	//若函数的形参为空接口，则实参可以为任意类型
	DoJob(myReader)
}
