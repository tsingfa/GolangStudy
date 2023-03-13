//对map遍历时删除元素能遍历到么？

package main

import (
	"fmt"
	"sync"
)

func main() {
	var m = map[int]int{1: 1, 2: 2, 3: 3} //only del key once, and not del the current iteration key
	var once sync.Once
	for i := range m {
		once.Do(func() {
			for _, key := range []int{1, 2, 3} {
				if key != i {
					fmt.Printf("当前 i:%d, del key:%d\n", i, key)
					delete(m, key)
					break
				}
			}
		})
		fmt.Printf("%d%d\n", i, m[i])
	}
}
