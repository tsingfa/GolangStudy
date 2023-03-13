package main

import "fmt"

func main() {
	//有可能遍历到，也有可能没遍历到新增元素
	//因为map的hash表初始化时会随机一个遍历开始的位置
	var m = map[int]int{1: 1, 2: 2, 3: 3}
	for i, _ := range m {
		m[4] = 4
		fmt.Printf("%d,%d\n", i, m[i])
	}
}
