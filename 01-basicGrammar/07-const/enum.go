package main

import "unsafe"

const ( //常量用作枚举
	//常量可以使用表达式的值
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a) //字符串对应结构体大小为16
)

/*
	typedef struct {
	char* buffer;    // 指向字符串的指针，大小8字节
	size_tlen;       // 字符串的长度，大小8字节
	} string;
*/

func main() {
	println(a, b, c) //abc 3 16
}
