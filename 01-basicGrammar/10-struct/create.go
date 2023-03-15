package main

import "fmt"

type Student struct {
	ID    int
	Name  string
	Age   int
	Score int
}

func main() {
	st1 := Student{ //键值对初始化
		ID:    100,
		Name:  "zhangsan",
		Age:   18,
		Score: 98,
	}
	st2 := Student{ //值列表初始化
		101,
		"lisi",
		20,
		97,
	}
	fmt.Printf("学生st1: %v\n", st1)
	fmt.Printf("学生st2: %v\n", st2)

	//结构体成员访问
	fmt.Printf("学生1的姓名是: %s\n", st1.Name)
	fmt.Printf("学生2的分数是: %d\n", st2.Score)
}
