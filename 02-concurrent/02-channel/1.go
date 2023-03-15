package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	ch <- 1
	close(ch)
	go func() {
		for i := 0; i < 5; i++ {
			v := <-ch
			fmt.Printf("v=%d\n", v)
		}
	}()
	time.Sleep(2 * time.Second)
}
