package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//假设值都为1，这里只赋值3个
	var arr = [2 * 1e7]int{1, 1, 1}

	start := time.Now()
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)

	for i, _ := range &arr {
		arr[i] = 0 //重置的效率很高
		//因为go对这种重置元素值为默认值的遍历有优化
	}

	runtime.ReadMemStats(&m2)
	fmt.Printf("代码块使用的内存大小为"+
		"：%v bytes\n", m2.Alloc-m1.Alloc)
	elapsed := time.Since(start)
	fmt.Printf("程序运行时间为：%s\n", elapsed)
}
