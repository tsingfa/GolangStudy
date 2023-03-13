package main

import "fmt"

func main() {
	v := []int{1, 2, 3}
	for i := range v {
		//循环会停止吗？
		//会！因为所遍历的是对目标做的拷贝（一个临时的中间体），它不变
		v = append(v, i)
		fmt.Printf("i=%v,v=%v\n", i, v)
	}
}
