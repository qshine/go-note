package main

import (
	"fmt"
	"os"
)

func main() {
	var f, err = os.Open("slice.go")
	if err != nil {
		fmt.Println("open error")
		return
	}
	defer f.Close()

	// 存储文件内容
	var content = []byte{}
	// 一次读取100字节, 注意不能是 make([]byte, 0, 100) 否则会出现死循环
	var buf = make([]byte, 100)

	for {
		// Read传入读取的字节数, 是切片的长度不是容量
		n, err := f.Read(buf)
		if n > 0 {
			// ...表示添加多个元素, 类似python中的 extend.
			// 此处是把一个切片放到另一个切片中
			content = append(content, buf[:n]...)

		}
		// 如果读取结束或其它错误退出
		if err != nil {
			break
		}

	}

	fmt.Println(string(content))

}
