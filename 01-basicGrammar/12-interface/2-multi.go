//go:build ignore
// +build ignore

/*
一种类型实现多个接口
*/
package main

import (
	"fmt"
)

type MyWriter interface {
	MyWriter(note string)
}

type MyReader interface {
	MyReader()
}

type ReadWriter struct {
}

func (wr ReadWriter) MyWriter(note string) {
	fmt.Println("Call ReadWriter MyWriter, Note =", note)
}

func (wr ReadWriter) MyReader() {
	fmt.Println("Call ReadWriter MyReader")
}

func main() {
	var readWriter ReadWriter
	readWriter.MyWriter("Hello Golang")
	readWriter.MyReader()
}
