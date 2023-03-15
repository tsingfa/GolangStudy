//go:build ignore
// +build ignore

/*
方法的定义与调用
*/
package main

import "fmt"

type Student struct {
	ID    int
	Name  string
	Age   int
	Score int
}

func (st *Student) GetName() string { //指定了接收者的函数
	return st.Name
}

func (st *Student) SetScore(score int) {
	st.Score = score
}

func (st Student) GetScore() int {
	return st.Score
}

func main() {
	st := &Student{
		ID:    100,
		Name:  "zhangsan",
		Age:   18,
		Score: 98,
	}
	fmt.Println("st指针:", st)
	fmt.Printf("st指针指向的地址:%p\n", st)  //0xc0000bc480
	fmt.Printf("st指针自身的地址:%p\n", &st) //0xc0000ca018 等价于 fmt.Println("st指针自身的地址:", &st)
	fmt.Printf("st指针指向的字段值:%v\n", *st)
	fmt.Printf("学生st的姓名是: %s\n", st.GetName())      // 调用Student的方法GetName
	fmt.Printf("设置前，学生st的分数是: %d\n", st.GetScore()) // 通过指针调用定义在【值类型】的方法GetScore
	st.SetScore(100)                                // 通过指针调用定义在【指针类型】上的方法
	fmt.Printf("设置后，学生st的分数是: %d\n", st.GetScore()) // 通过指针调用定义在【值类型】的方法GetScore
}
