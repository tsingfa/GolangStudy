// 测试对map的遍历、索引等
/*
	在Go语言中，`map`的键值映射是无序的。

	这是因为`map`是使用哈希表来实现的，
	哈希表中的键值对是根据键的哈希值来存储和查找的，
	而哈希值是随机的，因此键值对的顺序是不确定的。

	如果需要对键值对进行排序，
	可以将键值对存储在`slice`中，
	然后使用`sort`包中的函数进行排序。

	另外，从Go 1.12版本开始，`map`的遍历顺序是随机的，
	这是为了增加安全性和减少攻击面，
	防止攻击者利用`map`的遍历顺序来进行攻击。
*/

package main

import "fmt"

func main() {
	//`map`的键是唯一的，每个键只能对应一个值。
	//但是，不同的键可以对应相同的值。
	//key:value
	m := map[string]int{"c": 8, "d": 7, "x": 6, "y": 6}
	fmt.Println("m['c']=", m["c"])             //由键查找值，直接索引
	fmt.Println("m['nowhere']=", m["nowhere"]) //`key`不存在于`map`中，`value`将会是该值类型的零值。

	//若要由值查找键，则需遍历整个map，逐个比较键的值
	searchValue := 6
	for key, value := range m {
		if value == searchValue {
			fmt.Printf("key[%v]:value[%v]\n", key, value)
		}
	}

	//验证语句（注意与上面range循环的语句区分）
	value, ok := m["x"]
	fmt.Printf("value:%v,ok:%v\n", value, ok)
}
