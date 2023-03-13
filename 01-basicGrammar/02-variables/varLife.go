package main

import "fmt"

var globalStr string
var globalInt int

func main() {
	var localStr string
	var localInt int
	localStr = "first local"
	localInt = 2021
	globalInt = 1024
	globalStr = "first global"
	fmt.Printf("globalStr is %s\n", globalStr)   //globalStr is first global
	fmt.Printf("globalStr is %d\n", globalInt)   //globalStr is 1024
	fmt.Printf("localInt is %s\n", localStr)     //localInt is first local
	fmt.Printf("localInt int is %d\n", localInt) //localInt int is 2021
}
