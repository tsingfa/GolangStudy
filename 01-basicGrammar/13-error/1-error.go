package main

import (
	"errors"
	"fmt"
)

// 定义一个正数自加的函数，当传入的整数小于等于0的时候报错
func getPositiveSelfAdd(num int) (int, error) {
	if num <= 0 {
		return -1, fmt.Errorf("num is not a positive number")
	}
	return num + 1, nil
}

func main() {
	num1, err1 := getPositiveSelfAdd(1)
	fmt.Printf("nums is %d, err is %v\n", num1, err1)

	num2, err2 := getPositiveSelfAdd(-2)
	fmt.Printf("nums is %d, err is %v\n", num2, err2)

	err3 := errors.New("hello")
	err4 := errors.New("hello")
	fmt.Println(err3 == err4) //false
	//通过Error()方法拿到其中的error字符串信息，比较字符串
	fmt.Println(err3.Error() == err4.Error()) //true
}
