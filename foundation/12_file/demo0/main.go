package main

// 读取文件

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 一次全部读取
func test0() {
	file, err := os.Open("./unittest.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("open error")
		return
	}
	defer file.Close()

	// 定义字节数组
	var tmp = make([]byte, 128)
	n, err := file.Read(tmp)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))
}

// 循环读取
func test1() {
	file, err := os.Open("./unittest.txt")
	if err != nil {
		fmt.Println("open error")
		return
	}
	defer file.Close()

	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		// 把分批读取的内容存储到content中
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}

// 使用ioutil包读取
func test2() {
	content, err := ioutil.ReadFile("./unittest.txt")
	if err != nil {
		fmt.Println("open error")
		return
	}
	fmt.Println(string(content))
}

func main() {
	test0()
	fmt.Println("====")
	test1()
	fmt.Println("====")
	test2()
}
