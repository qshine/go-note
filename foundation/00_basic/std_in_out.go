package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// 创建一个buffer, 将字符串写入buffer中
	var b bytes.Buffer
	// 使用Write方法将字符串写入
	b.Write([]byte("hello "))

	// 将一个字符串拼接到Buffer里
	fmt.Fprintf(&b, "world")

	b.WriteTo(os.Stdout)

	// hello world
}
