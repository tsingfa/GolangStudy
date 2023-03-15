package main

import (
	"fmt"
	"time"
)

func myGroutine2(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v:myGroutine %s\n", i, name)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	go myGroutine2("goroutine1")
	go myGroutine2("goroutine2")
	time.Sleep(time.Second)
}
