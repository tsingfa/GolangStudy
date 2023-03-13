package main

import "fmt"

// "<<"左移运算，左移n位即为*(2^n)
// 如：3=(11)_2,左移2位为(1100)_2=12=3*(2^2)=3<<2,
const (
	i = 1 << iota //1*(2^0)=1
	j = 3 << iota //3*(2^1)=6
	k             //3*(2^2)=12
	l             //3*(2^3)=24
)

func main() {
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}
