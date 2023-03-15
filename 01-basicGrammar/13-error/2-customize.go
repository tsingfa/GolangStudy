/*
内置error只返回了错误信息，类似于error的message，
但是很多时候我们业务上还需要错误码，即error code。
因此，我们可以选择自定义error对象
*/
package main

import "fmt"

/*
//error的接口定义
type error interface {
   Error() string
}
*/

type MyError struct {
	code int
	msg  string
}

func (m MyError) Error() string {
	//MyError类型实现了error接口（可以当做error使用）
	return fmt.Sprintf("code:%d,msg:%v", m.code, m.msg)
}

func NewError(code int, msg string) error {
	return MyError{
		code: code,
		msg:  msg,
	}
}

func Code(err error) int {
	if e, ok := err.(MyError); ok {
		return e.code
	}
	return -1
}

func Msg(err error) string {
	if e, ok := err.(MyError); ok { //使用了error接口的断言
		return e.msg
	}
	return ""
}

func main() {
	err := NewError(100, "test MyError")
	fmt.Printf("code is %d, msg is %s", Code(err), Msg(err))
}
