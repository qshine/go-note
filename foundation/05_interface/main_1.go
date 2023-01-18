package main

/*
定义并实现一个interface
*/

import (
	"fmt"
)

// 定义一个Usb接口
type Usb interface {
	Connect()
}

type Pc struct {
	name string
}

func (p Pc) Connect() {
	fmt.Println(p.name)

}

func test_1() {
	// 声明a实现了Usb接口
	var a Usb
	a = Pc{
		name: "mac",
	}
	a.Connect()
	/*
		mac
	*/
}

func main() {
	test_1()
}
