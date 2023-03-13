// for-range 遍历取址的误用问题
package main

import "fmt"

func main() {
	arr := [2]int{1, 2}
	res := []*int{}
	//for-range 遍历取址的误用（该位置的v每次的地址相同，导致误用）
	for ind, v := range arr {
		fmt.Printf("&v:%v,v:%v\n", &v, v)
		res = append(res, &v)
		fmt.Printf("res[ind]:%v,*res[ind]:%v\n", res[ind], *res[ind])
	}

	fmt.Println(res[0], res[1])   //0xc00000e0b8 0xc00000e0b8
	fmt.Println(*res[0], *res[1]) //expect: 1 2	,but: 2 2
}
