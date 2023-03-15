package main

import "fmt"

func main() {
	res := deferRun3_1()
	fmt.Println(res) //2

	res = deferRun3_2()
	fmt.Println(res) //1

	res = deferRun3_3()
	fmt.Println(res) //1

	res = deferRun3_4()
	fmt.Println(res) //1
}

func deferRun3_1() (res int) { //返回值名称为"res"
	num := 1
	defer func() {
		res++
	}()
	/*	//return步骤
		1. 设置返回值（将result值设为num，即res=num=1）
		2. 执行defer语句（res++）
		3. 将结果返回（返回res=2）
	*/
	return num //实际返回的值是res
}
func deferRun3_2() int {
	var num int
	defer func() {
		num++
	}()
	return 1 //res=1
}
func deferRun3_3() int {
	num := 1
	defer func() {
		num++
	}()
	return num //res=1,num=2
}
func deferRun3_4() (res int) {
	num := 1
	defer func() {
		num++
	}()
	return num //res=1,num=2
}
