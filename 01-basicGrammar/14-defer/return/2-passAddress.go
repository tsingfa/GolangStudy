package main

import "fmt"

func main() {
	deferRun2()
}

func deferRun2() {
	var arr = [4]int{1, 2, 3, 4}
	defer printArr(&arr) //传进去的参数不变（这里是地址不变）

	arr[0] = 100
	return
}

func printArr(arr *[4]int) {
	for i := range arr {
		fmt.Println(arr[i])
	}
}
