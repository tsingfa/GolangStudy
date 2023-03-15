// 函数变量
package main

import (
	"fmt"
	"math"
)

func main() {
	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x) // 求一个数的平方根
	}

	/* 使用函数变量调用函数 */
	fmt.Println(getSquareRoot(9))
}
