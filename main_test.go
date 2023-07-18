package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestSign(t *testing.T) {

}

func TestIO(t *testing.T) {
	// 打开源文件（也可以是其他实现了 io.Reader 接口的对象）
	srcFile, err := os.Open("source.txt")
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	// 创建一个字节缓冲区，用于记录读取的数据
	var buf bytes.Buffer

	// 创建一个 io.TeeReader，它会同时读取源文件内容，并将内容写入字节缓冲区
	io.TeeReader(srcFile, &buf)

	// 使用 teeReader 处理读取的数据
	// ...

	// 在 buf 中记录的数据是和源文件内容一致的，你可以查看或处理 buf 中的数据
	fmt.Println("Recorded data:", buf.String())

	// 原始的 io.Reader（srcFile）保持不变，你可以继续使用它读取文件内容
}
