package main

import "fmt"

func main() {
	arr := [2]int{1, 2}
	res := []*int{}
	//for-range 遍历取址的误用（已修正）
	for ind, v := range arr {
		v1 := v
		fmt.Printf("&v:%v,v:%v\n", &v, v)     //v的地址不变
		fmt.Printf("&v1:%v,v1:%v\n", &v1, v1) //v1的地址会变
		res = append(res, &v1)
		//res=append(res,&arr[ind])			//直接索引获取原来元素的位置亦可
		fmt.Printf("res[ind]:%v,*res[ind]:%v\n", res[ind], *res[ind])
	}

	fmt.Println(res[0], res[1])   //0xc00000e0d0 0xc00000e100
	fmt.Println(*res[0], *res[1]) //expect: 1 2
}
