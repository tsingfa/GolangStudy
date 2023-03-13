//大数组遍历的效率问题

/*
在Go语言中，对于大量元素的`slice`和`map`进行遍历不会有内存浪费问题，
这是因为在遍历时，Go语言使用了指针来引用元素，而不是复制元素本身。

对于`slice`，Go语言使用了一个指向底层数组的指针和一个长度来表示`slice`。
在遍历`slice`时，Go语言只需要将指针向后移动，而不需要复制每个元素
。因此，无论`slice`的长度有多大，遍历`slice`的内存使用量都是固定的。

对于`map`，Go语言使用了一个哈希表来存储键值对。
在遍历`map`时，Go语言只需要遍历哈希表中的每个桶，然后遍历每个桶中的键值对。
由于哈希表的大小是固定的，因此无论`map`中有多少个键值对，遍历`map`的内存使用量也是固定的。

需要注意的是，如果在遍历`slice`或`map`时对元素进行修改，可能会导致内存分配。
这是因为在修改元素时，Go语言需要将元素复制到一个新的内存位置，以保证原始数据的不可变性。
因此，如果需要修改`slice`或`map`中的元素，建议使用指针或索引来访问元素，而不是直接复制元素。
*/

package main

import (
	"testing"
)

func test(b *testing.B) {
	b.ReportAllocs()
	//假设值都为1，这里只赋值3个
	var arr = [2 * 1e7]int{1, 1, 1}

	for i, n := range arr {
		//可改成取值&arr或
		// 切片arr[:]切片结构体指针类型，指向原底层数组，返回一个新的指针
		_ = i
		_ = n
	}

}

/*
	start := time.Now()

	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)

	runtime.ReadMemStats(&m2)
	fmt.Printf("代码块使用的内存大小为"+
		"：%v bytes\n", m2.Alloc-m1.Alloc) //16013576 bytes
	elapsed := time.Since(start)
	fmt.Printf("程序运行时间为：%s\n", elapsed) //23.0228ms
*/
