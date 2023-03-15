/*
资源释放：使用defer延迟调用机制，避免遗漏相关资源的回收问题
常见如：网络连接，数据库连接，以及文件句柄的资源的释放
*/
package main

import (
	"io"
	"os"
)

func CopyFile(dstFile, srcFile string) (wr int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {
		return
	}
	defer src.Close() //延迟调用，如果下方报error,return也可以不遗漏回收src

	dst, err := os.Create(dstFile)
	if err != nil {
		return
	}
	defer dst.Close()

	wr, err = io.Copy(dst, src)
	return wr, err
}
