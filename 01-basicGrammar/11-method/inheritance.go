// 方法的继承（组合实现）
package main

import "fmt"

type People struct {
	Name string
	Age  int
}

func (p *People) GetName() string {
	return p.Name
}

type Student struct {
	ID     int
	Score  int
	People // 将People作为Student的一个属性，注意不要加类型，这就是隐式继承
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
		Score: 98,
		People: People{
			Name: "zhangsan",
			Age:  18,
		},
	}
	fmt.Printf("学生st的分数是: %d\n", st.GetScore()) // 通过指针调用定义在值类型的方法GetScore                                         // 通过指针调用定义在指针类型上的方法
	fmt.Printf("学生st的姓名是: %s\n", st.GetName())

}
