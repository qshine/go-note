package main

import "fmt"

/*
嵌入接口
*/

type Connecter interface {
	Connect()
}

type USB interface {
	Name() string // 表示该方法返回值是string类型
	Connecter     // 嵌入另一个接口, 类似实现了继承
}

type PC struct {
	name string
}

func (pc PC) Name() string {
	return pc.name
}

func (pc PC) Connect() {
	fmt.Println("Connect: ", pc.name)
}

func Disconnect(usb USB) {
	if pc, ok := usb.(PC); ok {
		fmt.Println(ok)
		fmt.Println("Disconnect: ", pc.name)
	} else {
		fmt.Println("Unknown device !")
	}
}

func main() {
	// 声明b实现了USB接口
	var b USB
	b = PC{"computer"}
	fmt.Println(b.Name())
	b.Connect()
	Disconnect(b)

	/*
		computer
		Connect:  computer
		true
		Disconnect:  computer
	*/
}
