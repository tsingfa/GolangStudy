// helloworld.go
package main

//import的路径 = go.mod下的module name + 包相对于go.mod的相对目录
import (
	mathClass "GolangStudy/01-basicGrammar/01-golangStructure/myMath"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(mathClass.Add(1, 1))
	fmt.Println(mathClass.Sub(1, 1))
}
